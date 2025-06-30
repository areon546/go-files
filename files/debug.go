package files

import "github.com/areon546/go-helpers/helpers"

var debugMode = false

func debugPrint(a ...any) {
	if debugMode {
		helpers.Print("fileIO")
		helpers.Print(a...)
	}
}

func SetDebug(newDB bool) {
	debugMode = newDB
}
