package messages

import (
	"regexp"
	"strings"
	"text/template"
	"time"

	"mailarc/internal/mimecontent"
)

func TemplFuncs() template.FuncMap {
	return template.FuncMap{
		"formatDate":  tmplFormatDate,
		"formatAddr":  tmplFormatAddr,
		"decode":      tmplDecode,
		"decodeSlice": tmplFuncDecodeSlice,
		"isText":      tmplFuncIsText,
		"isImage":     tmplFuncIsImage,
	}
}

var addrSuffixPattern = regexp.MustCompile(`\s+\<.+?\>$`)

func tmplFormatDate(t time.Time) string {
	// TODO: To local time?
	return t.Format("Mon, 02 Jan 2006 15:04")
}

func tmplFormatAddr(s string) string {
	s = addrSuffixPattern.ReplaceAllString(s, "")
	s = strings.Trim(s, "\"")
	return s
}
func tmplDecode(s string) string {
	v, err := mimecontent.DecodeHeader(s)
	if err != nil {
		return s
	}
	return v
}

func tmplFuncDecodeSlice(a []string) string {
	b := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = tmplDecode(a[i])
	}
	return strings.Join(b, "<br />\n")
}

func tmplFuncIsText(content *mimecontent.Content) bool {
	return strings.HasPrefix(content.ContentType(), "text/")
}

func tmplFuncIsImage(content *mimecontent.Content) bool {
	return strings.HasPrefix(content.ContentType(), "image/")
}
