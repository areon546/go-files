package log

import (
	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

type (
	LogOutput interface {
		Output(a ...any)
	}

	PrintLogger struct{}
	fileLogger  struct{}
	NoLogger    struct{}

	Logger struct {
		logOut LogOutput
	}
)

func (logger PrintLogger) Output(a ...any) {
	helpers.Print(a...)
}

func (logger fileLogger) Output(a ...any) {
	files.OpenFile("doesntExist.log")
}

func (logger NoLogger) Output(a ...any) {
}

func (logger Logger) Log(a ...any) {
	logger.logOut.Output(a...)
}
