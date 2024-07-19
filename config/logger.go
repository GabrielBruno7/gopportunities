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

func (log *Logger) Debug(value...interface{}) {
	log.debug.Println(value...)
}

func (log *Logger) Info(value...interface{}) {
	log.info.Println(value...)
}

func (log *Logger) Warning(value...interface{}) {
	log.warning.Println(value...)
}

func (log *Logger) Error(value...interface{}) {
	log.err.Println(value...)
}


func (log *Logger) DebugFormated(format string, value ...interface{}) {
	log.debug.Printf(format, value...)
}

func (log *Logger) InfoFormated(format string, value ...interface{}) {
	log.info.Printf(format, value...)
}

func (log *Logger) WarningFormated(format string, value ...interface{}) {
	log.warning.Printf(format, value...)
}

func (log *Logger) ErrorFormated(format string, value ...interface{}) {
	log.err.Printf(format, value...)
}