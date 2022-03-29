package logger

import (
	"io"
	"log"
)

func NewLogger(out io.Writer, prefix string) *log.Logger {
	prefix += " "

	return log.New(out, prefix, log.LstdFlags)
}
