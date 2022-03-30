package logger

import (
	"io"
	"log"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger(out io.Writer, prefix string) *Logger {
	prefix += " "

	return &Logger{
		logger: log.New(out, prefix, log.LstdFlags),
	}
}

func (l Logger) Println(data any) {
	l.logger.Println(data)
}
