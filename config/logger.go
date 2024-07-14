package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate)

	return &Logger{
		debug: log.New(writer, "Debug: ", logger.Flags()),
		info: log.New(writer, "Info: ", logger.Flags()),
		warning: log.New(writer, "Warning: ", logger.Flags()),
		err: log.New(writer, "Error: ", logger.Flags()),
		writer: writer,
	}
}
