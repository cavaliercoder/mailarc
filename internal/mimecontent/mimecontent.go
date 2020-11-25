package mimecontent

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/textproto"
	"path/filepath"
	"strconv"
	"strings"
)

type Content struct {
	Path     string
	Headers  textproto.MIMEHeader
	Body     []byte
	Children []*Content
}

func (c *Content) String() string                  { return string(c.Body) }
func (c *Content) ContentDisposition() string      { return c.Headers.Get("Content-Disposition") }
func (c *Content) ContentType() string             { return c.Headers.Get("Content-Type") }
func (c *Content) ContentTransferEncoding() string { return c.Headers.Get("Content-Transfer-Encoding") }

func (c *Content) Root() *Content {
	if len(c.Children) == 0 {
		return c
	}
	switch c.MediaType() {
	case "multipart/mixed":
		return c.Children[0]
	case "multipart/alternative":
		return c.Children[len(c.Children)-1]
	case "multipart/related":
		// TODO: honor the start param
		return c.Children[0]
	}
	return c
}

func (c *Content) MediaType() string {
	s, _, err := mime.ParseMediaType(c.ContentType())
	if err != nil {
		return ""
	}
	return s
}

func (c *Content) IsAttachment() bool {
	disposition, _, err := mime.ParseMediaType(c.ContentDisposition())
	if err != nil {
		return false
	}
	return disposition == "attachment"
}

func (c *Content) IsMultipart() bool { return strings.HasPrefix(c.MediaType(), "multipart/") }

func (c *Content) Filename() string {
	disposition, params, err := mime.ParseMediaType(c.ContentDisposition())
	if err != nil {
		return ""
	}
	if disposition == "attachment" {
		if s := params["filename"]; s != "" {
			return s
		}
	}
	return ""
}

func parsePath(path string) []int {
	parts := strings.Split(path, "/")
	a := make([]int, len(parts))
	for i, part := range parts {
		v, err := strconv.Atoi(part)
		if err != nil {
			return nil
		}
		a[i] = v
	}
	return a
}

func (c *Content) GetPart(path string) *Content {
	if path == "/" {
		return c
	}
	parts := parsePath(path)
	if parts == nil {
		return nil
	}
	for len(parts) > 0 {
		i := parts[0]
		parts = parts[1:]
		if i >= len(c.Children) {
			return nil
		}
		c = c.Children[i]
	}
	return c
}

func readBody(r io.Reader, h textproto.MIMEHeader, path string) (c *Content, err error) {
	c = &Content{
		Path:     path,
		Headers:  h,
		Children: make([]*Content, 0),
	}

	// decode media type
	contentType := h.Get("Content-Type")
	if contentType == "" {
		contentType = "message/rfc822"
	}
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, err
	}

	// read single-part body
	if !strings.HasPrefix(mediaType, "multipart/") {
		switch h.Get("Content-Transfer-Encoding") {
		case "base64":
			r = base64.NewDecoder(base64.StdEncoding, r)
		}
		c.Body, err = ioutil.ReadAll(r)
		return
	}

	// read multipart children
	var child *Content
	var p *multipart.Part
	if !strings.HasPrefix(mediaType, "multipart/") {
		return
	}
	mr := multipart.NewReader(r, params["boundary"])
	for i := 0; ; i++ {
		p, err = mr.NextPart()
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return
		}
		child, err = readBody(p, p.Header, filepath.Join(path, fmt.Sprint(i)))
		if err == io.ErrUnexpectedEOF {
			// this can happen when a closing multi-part boundary is omitted
			err = nil
			c.Children = append(c.Children, child)
			return
		}
		if err != nil {
			return
		}
		c.Children = append(c.Children, child)
	}
	return
}

func Read(r *bufio.Reader) (*Content, error) {
	tpr := textproto.NewReader(r)
	h, err := tpr.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}
	return readBody(r, h, "/")
}

func DecodeHeader(s string) (string, error) {
	var err error
	dec := new(mime.WordDecoder)
	s, err = dec.DecodeHeader(s)
	if err != nil {
		return "", err
	}
	return s, nil
}
