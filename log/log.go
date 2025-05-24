package log

import (
	"fmt"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

type (
	LogOutput interface {
		Output(a ...any)
	}

	printLogger struct{}
	fileLogger  struct {
		logFile files.TextFile
	}
	noLogger struct{}

	Logger struct {
		logOut LogOutput
	}
)

func NewPrintLogger() *printLogger {
	return &printLogger{}
}

func (logger printLogger) Output(a ...any) {
	helpers.Print(a...)
}

func NewFileLogger(filePath string) *fileLogger {
	return &fileLogger{logFile: *files.NewTextFile(filePath)}
}

func (logger fileLogger) Output(a ...any) {
	logger.logFile.AppendNewLine(fmt.Sprint(a...))
}

func NewNoLogger() *noLogger {
	return &noLogger{}
}

func (logger noLogger) Output(a ...any) {
}

// Logger

func NewLogger(logger *LogOutput) *Logger {
	return &Logger{logOut: *logger}
}

func (logger Logger) Log(a ...any) {
	logger.logOut.Output(a...)
}
