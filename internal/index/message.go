package index

import (
	"mailarc/internal/mimecontent"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/analysis/lang/en"
	"github.com/blevesearch/bleve/mapping"
)

const MessageEntryType = "Message"

const (
	MessageFieldDate    = "Date"
	MessageFieldSubject = "Subject"
	MessageFieldFrom    = "From"
	MessageFieldTo      = "To"
	MessageFieldContent = "Content"
)

var MessageFields = []string{
	MessageFieldDate,
	MessageFieldSubject,
	MessageFieldFrom,
	MessageFieldTo,
	MessageFieldContent,
}

type MessageEntry struct {
	UID     string
	Date    time.Time
	Subject string
	From    string
	To      string

	content []string
}

func (c *MessageEntry) Type() string { return MessageEntryType }

func newMessageEntry(uid string, content *mimecontent.Content) *MessageEntry {
	// parse message date
	date, err := mimecontent.ParseDateTime(content.Headers.Get("Date"))
	if err != nil {
		date = time.Time{}
	}

	m := &MessageEntry{
		UID:     uid,
		Date:    date,
		Subject: content.Headers.Get(MessageFieldSubject),
		From:    content.Headers.Get(MessageFieldFrom),
		To:      content.Headers.Get(MessageFieldTo),
		content: make([]string, 0),
	}

	// extract body and attachent content
	queue := make([]*mimecontent.Content, 1)
	queue[0] = content
	for len(queue) > 0 {
		part := queue[0]
		queue = queue[1:]

		switch part.MediaType() {
		case "text/html":
		case "text/plain":
			// TODO: limit size?
			m.content = append(m.content, string(part.Body))
		default:
			// ignore
		}
		for _, child := range part.Children {
			queue = append(queue, child)
		}
	}
	return m
}

func GetMessageMapping() (doctype string, dm *mapping.DocumentMapping) {
	dm = bleve.NewDocumentMapping()

	date := bleve.NewDateTimeFieldMapping()
	date.Store = true

	subject := bleve.NewTextFieldMapping()
	subject.Analyzer = en.AnalyzerName
	subject.Store = true

	addr := bleve.NewTextFieldMapping()
	addr.Analyzer = keyword.Name
	addr.Store = true

	content := bleve.NewTextFieldMapping()
	content.Analyzer = en.AnalyzerName
	content.Store = false

	dm.AddFieldMappingsAt(MessageFieldSubject, subject)
	dm.AddFieldMappingsAt(MessageFieldFrom, addr)
	dm.AddFieldMappingsAt(MessageFieldTo, addr)
	dm.AddFieldMappingsAt(MessageFieldDate, date)
	dm.AddFieldMappingsAt(MessageFieldContent, content)
	return MessageEntryType, dm
}
