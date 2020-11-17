package store

import (
	"compress/gzip"
	"io"
)

type reader struct {
	r   io.ReadCloser
	gzr *gzip.Reader
}

func newReader(r io.ReadCloser) (io.ReadCloser, error) {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}
	return &reader{r: r, gzr: gzr}, nil
}

func (c *reader) Read(p []byte) (n int, err error) {
	return c.gzr.Read(p)
}

func (c *reader) Close() (err error) {
	if err = c.gzr.Close(); err != nil {
		return
	}
	return c.r.Close()
}
