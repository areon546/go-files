package files

import (
	"reflect"
	"strings"

	"github.com/areon546/go-helpers/helpers"
)

// helper functions

func handle(err error) {
	// helpers.Handle(err)
	helpers.Handle(err)
}

func print(a ...any) {
	helpers.Print(a...)
}

func printf(s string, a ...any) {
	helpers.Printf(s, a...)
}

func search(s string, ss []string) int { return helpers.Search(s, ss) }

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func bytesToString(b []byte) string { return helpers.BytesToString(b) }

func FilesEqual(a, b File) bool {
	return reflect.DeepEqual(a, b)
}

func trimFiletype(filename, filetype string) string {
	if filenameContains(filename, filetype) {
		return removeFiletype(filename, filetype)
	}
	return filename
}

func filenameContains(filename, filetype string) bool {
	filenameComponents := strings.Split(filename, ".")
	if reflect.DeepEqual(filenameComponents[len(filenameComponents)-1], filetype) {
		// filename type is filetype specified
		return true
	}

	return false
}

func removeFiletype(filename, filetype string) string {
	return strings.TrimRight(filename, ("." + filetype))
}
