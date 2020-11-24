package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var logLevelDebug = os.Getenv("LOGLEVEL") == "DEBUG"

func Checksum(b []byte) string {
	digest := sha256.Sum256(b)
	s := hex.EncodeToString(digest[:])
	return s[:40]
}

func ResolveTilde(s string) string {
	if s[0] != '~' {
		return s
	}
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, s[1:])
}

func Logf(format string, a ...interface{}) {
	log.Printf(format, a...)
}

func LogDebugf(format string, a ...interface{}) {
	if !logLevelDebug {
		return
	}
	s := fmt.Sprintf(format, a...)
	Logf("DEBUG: %s", s)
}
