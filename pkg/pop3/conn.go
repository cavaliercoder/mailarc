package pop3

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"os"
	"strings"
)

var (
	crlf = []byte{'\r', '\n'}
)

type Conn interface {
	net.Conn

	ReadLine() (b []byte, err error)
	WriteLine(format string, a ...interface{}) (err error)
}

type pop3Conn struct {
	net.Conn
	textConn *textproto.Conn

	traceWriter io.Writer
}

func NewConn(conn net.Conn) Conn {
	var w io.Writer
	if os.Getenv("DEBUG") != "" {
		w = os.Stderr
	}
	return &pop3Conn{
		Conn:        conn,
		textConn:    textproto.NewConn(conn),
		traceWriter: w,
	}
}

func Dial(addr string) (Conn, error) {
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		return nil, err
	}
	return NewConn(conn), nil
}

func DialWithTLS(addr string, config *tls.Config) (Conn, error) {
	conn, err := tls.Dial("tcp4", addr, config)
	if err != nil {
		return nil, err
	}
	return NewConn(conn), nil
}

func (c *pop3Conn) ReadLine() (b []byte, err error) {
	b, err = c.textConn.ReadLineBytes()
	if err != nil {
		return
	}
	if c.traceWriter != nil {
		c.traceWriter.Write(b)
		c.traceWriter.Write([]byte{'\n'})
	}
	return
}

func (c *pop3Conn) WriteLine(format string, a ...interface{}) (err error) {
	if _, err = fmt.Fprintf(c.textConn.W, format, a...); err != nil {
		return
	}
	if _, err = c.textConn.W.Write(crlf); err != nil {
		return
	}
	if err = c.textConn.W.Flush(); err != nil {
		return
	}
	if c.traceWriter != nil {
		s := fmt.Sprintf(format, a...)
		if strings.HasPrefix(s, "PASS ") {
			s = "PASS ********"
		}
		fmt.Fprintf(c.traceWriter, "> %s\n", s)
	}
	return
}

func (c *pop3Conn) Close() error { return c.textConn.Close() }
