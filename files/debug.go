package files

import "github.com/areon546/go-helpers/helpers"

var debugMode = false

func debugPrint(a ...any) {
	if debugMode {
		helpers.Print(a...)
	}
}

func ToggleDebugPrint() {
	debugMode = !debugMode
}
