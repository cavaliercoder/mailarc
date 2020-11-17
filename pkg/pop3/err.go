package pop3

import "fmt"

type POP3Error error

func newError(format string, a ...interface{}) error {
	s := fmt.Sprintf(format, a...)
	return POP3Error(fmt.Errorf("pop3: %s", s))
}
