package log

import "github.com/areon546/go-helpers/helpers"

var debugMode = false

func DebugPrint(a ...any) {
	if debugMode {
		helpers.Print(a...)
	}
}

func SetDebug(newDB bool) {
	debugMode = newDB
}
