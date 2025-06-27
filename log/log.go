package log

import (
	"fmt"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

type (
	LogOutput interface {
		Output(a ...any) // used to transfer data to any buffer
		Close()          // used to close any IO streams that need to be closed, eg File IO
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

// PRINT LOGGER
func NewPrintLogger() *printLogger {
	return &printLogger{}
}

func (logger printLogger) Output(a ...any) {
	helpers.Print(a...)
}

func (logger printLogger) Close() {
	DebugPrint("closing print logger: ")
}

// FILE LOGGER
func NewFileLogger(filePath string) *fileLogger {
	return &fileLogger{logFile: *files.NewTextFile(filePath)}
}

func (logger fileLogger) Output(a ...any) {
	DebugPrint("logfile contents: ", logger.logFile.Contents())

	logger.logFile.AppendNewLine(fmt.Sprint(a...))
}

func (logger fileLogger) Close() {
	DebugPrint("closing file logger ")

	logger.logFile.Close()
}

// NO LOG
func NewNoLogger() *noLogger {
	return &noLogger{}
}

func (logger noLogger) Output(a ...any) {
	DebugPrint("no log ", a)
}

func (logger noLogger) Close() {
	DebugPrint("closing no loger")
}

// Logger

func NewLogger(logger *LogOutput) *Logger {
	return &Logger{logOut: *logger}
}

func (logger Logger) Log(a ...any) {
	logger.logOut.Output(a...)
}
