package store

import (
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cavaliercoder/mailarc/internal/util"
)

var (
	ErrNotFound = errors.New("store: message not found")
)

type ReadStore interface {
	List() ([]string, error)
	Exists(name string) (bool, error)
	Open(name string) (r io.ReadCloser, err error)
}

type WriteStore interface {
	Store(name string, r io.Reader) (err error)
	Delete(name string) (err error)
}

type Store interface {
	ReadStore
	WriteStore
}

type store struct {
	path string
}

func New(path string) Store {
	util.LogDebugf("Mailbox store: %s", path)
	return &store{path: path}
}

func (c *store) getFilePath(name string) string {
	return filepath.Join(c.path, fmt.Sprintf("%s.eml.gz", name))
}

func (c *store) Open(name string) (r io.ReadCloser, err error) {
	f, err := os.Open(c.getFilePath(name))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return newReader(f)
}

func (c *store) Store(name string, r io.Reader) (err error) {
	f, err := os.Create(c.getFilePath(name))
	if err != nil {
		return err
	}
	defer f.Close()
	gzw, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		return err
	}
	defer gzw.Close()
	_, err = io.Copy(gzw, r)
	return err
}

func (c *store) Exists(name string) (bool, error) {
	_, err := os.Stat(c.getFilePath(name))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (c *store) List() ([]string, error) {
	dir, err := ioutil.ReadDir(c.path)
	if err != nil {
		return nil, err
	}
	a := make([]string, 0, 64)
	for _, fi := range dir {
		name := fi.Name()
		if !strings.HasSuffix(name, ".eml.gz") {
			continue
		}
		name = name[:len(name)-7]
		a = append(a, name)
	}
	return a, nil
}

func (c *store) Delete(name string) (err error) {
	ok, err := c.Exists(name)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotFound
	}
	return os.Remove(c.getFilePath(name))
}

func Checksum(r io.Reader) (s string, err error) {
	h := sha256.New()
	if _, err = io.Copy(h, r); err != nil {
		return
	}
	digest := hex.EncodeToString(h.Sum(nil))
	return digest[:40], nil
}
