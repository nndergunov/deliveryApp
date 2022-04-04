package logger

import (
	"io"
	"log"
)

// Logger is the default app logger wrapper.
type Logger struct {
	logger *log.Logger
}

// NewLogger returns new instance of logger.Logger.
func NewLogger(out io.Writer, prefix string) *Logger {
	prefix += " "

	return &Logger{
		logger: log.New(out, prefix, log.LstdFlags),
	}
}

// Printf prints data to the specified Output in similar to fmt.Printf way.
func (l Logger) Printf(format string, data ...any) {
	l.logger.Printf(format, data...)
}

// Println prints data to the specified Output in similar to fmt.Println way.
func (l Logger) Println(data ...any) {
	l.logger.Println(data...)
}

// Print prints data to the specified Output in similar to fmt.Print way.
func (l Logger) Print(data ...any) {
	l.logger.Print(data...)
}

// Fatal acts similar to log.Fatal.
func (l Logger) Fatal(data ...any) {
	l.logger.Fatal(data...)
}

// Fatalf acts similar to log.Fatalf.
func (l Logger) Fatalf(format string, data ...any) {
	l.logger.Fatalf(format, data...)
}

// Fatalln acts similar to log.Fatalln.
func (l Logger) Fatalln(data ...any) {
	l.logger.Fatalln(data...)
}

// Panic acts similar to log.Panic.
func (l Logger) Panic(data ...any) {
	l.logger.Panic(data...)
}

// Panicf acts similar to log.Panicf.
func (l Logger) Panicf(format string, data ...any) {
	l.logger.Panicf(format, data...)
}

// Panicln acts similar to log.Panicln.
func (l Logger) Panicln(data ...any) {
	l.logger.Panicln(data...)
}
