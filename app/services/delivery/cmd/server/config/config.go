package config

import (
	"time"
)

type Config struct {
	Address           string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	ReadHeaderTimeout time.Duration
}

func GetConfig() Config {
	var (
		readTime         = 5 * time.Second
		writeTime        = 5 * time.Second
		idleTime         = 300 * time.Second
		readerHeaderTime = 5 * time.Second
	)

	return Config{
		Address:           ":8080",
		ReadTimeout:       readTime,
		WriteTimeout:      writeTime,
		IdleTimeout:       idleTime,
		ReadHeaderTimeout: readerHeaderTime,
	}
}
