package messages

import (
	"mime"
	"strings"

	"github.com/cavaliercoder/mailarc/internal/mimecontent"
)

func tmplFuncDecode(s string) string {
	var err error
	dec := new(mime.WordDecoder)
	s, err = dec.DecodeHeader(s)
	if err != nil {
		panic(err) // TODO
	}
	return s
}

func tmplFuncDecodeSlice(a []string) string {
	b := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = tmplFuncDecode(a[i])
	}
	return strings.Join(b, "<br />\n")
}

func tmplFuncIsText(content *mimecontent.Content) bool {
	return strings.HasPrefix(content.ContentType(), "text/")
}

func tmplFuncIsImage(content *mimecontent.Content) bool {
	return strings.HasPrefix(content.ContentType(), "image/")
}
