package mimecontent

import (
	"fmt"
	"net/mail"
	"regexp"
	"time"

	"mailarc/internal/util"
)

var dateFormats = []string{
	// RFC 822-ish
	"2 Jan 06 15:04",
	"2 Jan 06 15:04 MST",
	"2 Jan 06 15:04 -0700",
	"2 Jan 06 15:04:05",
	"2 Jan 06 15:04:05 MST",
	"2 Jan 06 15:04:05 -0700",
	"2 Jan 2006 15:04:05",
	"2 Jan 2006 15:04:05 MST",
	"2 Jan 2006 15:04:05 -0700",

	// RFC 1123-ish
	"Mon, 2 Jan 2006 15:04:05 MST",
	"Mon, 2 Jan 2006 15:04:05 -0700",
	"Mon, 2 Jan 2006 15:04:05 -0700 (MST)",

	// RFC 3339-ish
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02T15:04:05Z",

	// ANSIC-ish
	"Mon Jan 2 15:04:05 2006",
}

var utPattern = regexp.MustCompile(`\bUT\b`)

func ParseDateTime(s string) (t time.Time, err error) {
	if s == "[date]" {
		//  TODO
		util.LogDebugf("WTF???? %s", s)
		return time.Time{}, nil
	}
	if s == "" {
		return time.Time{}, nil
	}

	// UT is a vaild RFC 822 timezone but not recognized by go
	s = utPattern.ReplaceAllString(s, "UTC")

	t, err = mail.ParseDate(s)
	if err == nil {
		return
	}
	for _, format := range dateFormats {
		t, err = time.Parse(format, s)
		if err == nil {
			return
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse date: %s", s)
}
