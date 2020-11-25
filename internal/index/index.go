package index

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/blevesearch/bleve"
	bleve_query "github.com/blevesearch/bleve/search/query"

	"mailarc/internal/mimecontent"
	"mailarc/internal/util"
)

type Index interface {
	Add(ctx context.Context, name string, content *mimecontent.Content) error
	Search(
		ctx context.Context,
		query string,
		offset int,
		size int,
	) ([]*MessageEntry, error)
	Close() error
}

type bleveIndex struct {
	index bleve.Index
}

func Open(path string) (Index, error) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return newIndex(path)
	}
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, errors.New("index path should be a directory")
	}
	index, err := bleve.Open(path)
	if err != nil {
		return nil, err
	}
	util.LogDebugf("Using index: %s", path)
	return &bleveIndex{index: index}, nil
}

func newIndex(path string) (Index, error) {
	mapping := bleve.NewIndexMapping()
	mapping.AddDocumentMapping(GetMessageMapping())
	index, err := bleve.New(path, mapping)
	if err != nil {
		return nil, err
	}
	util.LogDebugf("Created index: %s", path)
	return &bleveIndex{index: index}, nil
}

func (c *bleveIndex) Add(ctx context.Context, name string, content *mimecontent.Content) error {
	msg := newMessageEntry(name, content)
	return c.index.Index(name, msg)
}

func getStringField(fields map[string]interface{}, name string) (s string) {
	v, ok := fields[name]
	if !ok {
		return
	}
	s, ok = v.(string)
	if !ok {
		return fmt.Sprintf("%v", v)
	}
	return
}

func getDateField(fields map[string]interface{}, name string) (t time.Time) {
	s := getStringField(fields, name)
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return time.Time{}
	}
	return
}

func getQuery(s string) bleve_query.Query {
	if s == "" {
		return bleve.NewMatchAllQuery()
	}
	return bleve.NewMatchQuery(s)
}

func (c *bleveIndex) Search(
	ctx context.Context,
	query string,
	offset int,
	size int,
) ([]*MessageEntry, error) {
	q := getQuery(query)
	search := bleve.NewSearchRequestOptions(q, size, offset, false)
	search.Explain = true
	search.IncludeLocations = true
	search.Size = 10000
	search.Fields = MessageFields
	search.SortBy([]string{"-" + MessageFieldDate})
	results, err := c.index.SearchInContext(ctx, search)
	if err != nil {
		return nil, err
	}
	a := make([]*MessageEntry, len(results.Hits))
	for i, result := range results.Hits {
		a[i] = &MessageEntry{
			UID:     result.ID,
			Date:    getDateField(result.Fields, MessageFieldDate),
			Subject: getStringField(result.Fields, MessageFieldSubject),
			From:    getStringField(result.Fields, MessageFieldFrom),
			To:      getStringField(result.Fields, MessageFieldTo),
		}
	}
	return a, nil
}

func (c *bleveIndex) Close() error { return c.index.Close() }
