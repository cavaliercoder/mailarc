package pop3

import (
	"bytes"
)

var (
	responseOK  = []byte{'+', 'O', 'K'}
	responseErr = []byte{'-', 'E', 'R', 'R'}
)

type Session interface {
	Cmd(format string, a ...interface{}) (b []byte, err error)
	Stat() (count int, size int, err error)
	ListOne(msgNum int) (size int, err error)
	ListAll() (sizes map[int]int, err error)
	Retr(msgNum int) (b []byte, err error)
	Dele(msgNum int) (err error)
	Noop() (err error)
	Quit() (err error)
	UIDLOne(msgNum int) (uid string, err error)
	UIDLAll() (uids map[int]string, err error)
	Close() error
}

type session struct {
	conn Conn
}

func NewSession(conn Conn, username, password string) (Session, error) {
	// discard welcome message
	c := &session{conn: conn}
	if _, err := conn.ReadLine(); err != nil {
		return nil, err
	}
	// login
	if _, err := c.Cmd("USER %s", username); err != nil {
		return nil, err
	}
	if _, err := c.Cmd("PASS %s", password); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *session) Cmd(format string, a ...interface{}) (b []byte, err error) {
	if err = c.conn.WriteLine(format, a...); err != nil {
		return
	}
	b, err = c.conn.ReadLine()
	if err != nil {
		return
	}
	if len(b) >= 3 && bytes.Equal(b[:3], responseOK) {
		if len(b) > 4 {
			b = b[4:]
		} else {
			b = []byte{}
		}
		return
	}
	if len(b) >= 4 && bytes.Equal(b[:4], responseErr) {
		if len(b) > 5 {
			err = newError(string(b[5:]))
		} else {
			err = newError("unspecified server error")
		}
		return
	}
	return
}

func (c *session) Stat() (count int, size int, err error) {
	var b []byte
	b, err = c.Cmd("STAT")
	parts := bytes.Split(b, []byte{' '})
	if len(parts) != 2 {
		err = newError("unrecognized STAT response")
		return
	}
	if count, err = atoi(parts[0]); err != nil {
		err = newError("unrecognized STAT response")
		return
	}
	if size, err = atoi(parts[1]); err != nil {
		err = newError("unrecognized STAT response")
		return
	}
	return
}

func parseList(b []byte) (msgNum, size int, err error) {
	parts := bytes.Split(b, []byte{' '})
	if len(parts) != 2 {
		err = newError("unrecognized LIST response")
		return
	}
	if msgNum, err = atoi(parts[0]); err != nil {
		return
	}
	if size, err = atoi(parts[1]); err != nil {
		return
	}
	return
}

func (c *session) ListOne(msgNum int) (size int, err error) {
	var b []byte
	if b, err = c.Cmd("LIST %d", msgNum); err != nil {
		return
	}
	var n int
	if n, size, err = parseList(b); err != nil {
		return
	}
	if n != msgNum {
		err = newError("wrong message number returned")
	}
	return
}

func (c *session) ListAll() (sizes map[int]int, err error) {
	if _, err = c.Cmd("LIST"); err != nil {
		return
	}
	var msgNum, size int
	sizes = make(map[int]int)
	scanner := NewScanner(c.conn)
	for scanner.Scan() {
		if msgNum, size, err = parseList(scanner.Bytes()); err != nil {
			return
		}
		sizes[msgNum] = size
	}
	err = scanner.Err()
	return
}

func (c *session) Retr(msgNum int) (b []byte, err error) {
	if _, err = c.Cmd("RETR %d", msgNum); err != nil {
		return
	}
	buf := bytes.NewBuffer(nil)
	scanner := NewScanner(c.conn)
	for scanner.Scan() {
		buf.Write(scanner.Bytes())
		buf.Write([]byte{'\n'})
	}
	if err = scanner.Err(); err != nil {
		return
	}
	b = buf.Bytes()
	return
}

func (c *session) Dele(msgNum int) (err error) {
	_, err = c.Cmd("DELE %d", msgNum)
	return
}

func (c *session) Noop() (err error) {
	_, err = c.Cmd("NOOP")
	return
}

func (c *session) Quit() (err error) {
	_, err = c.Cmd("QUIT")
	return
}

func parseUIDL(b []byte) (msgNum int, uid string, err error) {
	parts := bytes.Split(b, []byte{' '})
	if len(parts) != 2 {
		err = newError("unrecognized UIDL response")
		return
	}
	if msgNum, err = atoi(parts[0]); err != nil {
		return
	}
	uid = string(parts[1])
	return
}

func (c *session) UIDLAll() (uids map[int]string, err error) {
	if _, err = c.Cmd("UIDL"); err != nil {
		return
	}
	var msgNum int
	var uid string
	uids = make(map[int]string)
	scanner := NewScanner(c.conn)
	for scanner.Scan() {
		if msgNum, uid, err = parseUIDL(scanner.Bytes()); err != nil {
			return
		}
		uids[msgNum] = uid
	}
	err = scanner.Err()
	return
}

func (c *session) UIDLOne(msgNum int) (uid string, err error) {
	var b []byte
	if b, err = c.Cmd("UIDL %d", msgNum); err != nil {
		return
	}
	var n int
	if n, uid, err = parseUIDL(b); err != nil {
		return
	}
	if n != msgNum {
		err = newError("wrong message number returned")

	}
	return
}

func (c *session) Close() (err error) {
	return c.conn.Close()
}

func atoi(b []byte) (v int, err error) {
	for i := 0; i < len(b); i++ {
		v *= 10
		if b[i] < '0' || b[i] > '9' {
			return 0, newError("error decoding integer")
		}
		v += int(b[i] - '0')
	}
	return
}
