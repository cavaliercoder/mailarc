package pop3

import "bytes"

type Scanner struct {
	conn Conn
	err  error
	b    []byte
	done bool
}

func NewScanner(conn Conn) *Scanner { return &Scanner{conn: conn} }

func (c *Scanner) Scan() bool {
	if c.done || c.conn == nil {
		return false
	}
	c.b, c.err = c.conn.ReadLine()
	if c.err != nil || bytes.Equal(c.b, []byte{'.'}) {
		c.done = true
		return false
	}
	return true
}

func (c *Scanner) Err() error { return c.err }

func (c *Scanner) Bytes() []byte {
	b := make([]byte, len(c.b))
	copy(b, c.b)
	return b
}

func (c *Scanner) Text() string {
	return string(c.b)
}
