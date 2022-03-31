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

func (l Logger) Printf(format string, data ...any) {
	l.logger.Printf(format, data...)
}

func (l Logger) Println(data ...any) {
	l.logger.Println(data...)
}

func (l Logger) Print(data ...any) {
	l.logger.Print(data...)
}

func (l Logger) Fatal(data ...any) {
	l.logger.Fatal(data...)
}

func (l Logger) Fatalf(format string, data ...any) {
	l.logger.Fatalf(format, data...)
}

func (l Logger) Fatalln(data ...any) {
	l.logger.Fatalln(data...)
}

func (l Logger) Panic(data ...any) {
	l.logger.Panic(data...)
}

func (l Logger) Panicf(format string, data ...any) {
	l.logger.Panicf(format, data...)
}

func (l Logger) Panicln(data ...any) {
	l.logger.Panicln(data...)
}
