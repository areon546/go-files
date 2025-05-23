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

	PrintLogger struct{}
	FileLogger  struct {
		logFile files.TextFile
	}
	NoLogger struct{}

	Logger struct {
		logOut LogOutput
	}
)

func newPrintLogger() *PrintLogger {
	return &PrintLogger{}
}

func (logger PrintLogger) Output(a ...any) {
	helpers.Print(a...)
}

func newFileLogger(filePath string) *FileLogger {
	return &FileLogger{logFile: *files.NewTextFile(filePath)}
}

func (logger FileLogger) Output(a ...any) {
	logger.logFile.AppendNewLine(fmt.Sprint(a...))
}

func NewNoLogger() *NoLogger {
	return &NoLogger{}
}

func (logger NoLogger) Output(a ...any) {
}

// Logger

func NewLogger(logger *LogOutput) *Logger {
	return &Logger{logOut: *logger}
}

func (logger Logger) Log(a ...any) {
	logger.logOut.Output(a...)
}
