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

func splitFileName(filename string) (name, suffix string) {
	stringSections := strings.Split(filename, ".")
	// print(stringSections)

	if len(stringSections) == 1 {
		name = stringSections[0]
		suffix = ""
		return
	}

	if len(stringSections) > 1 {
		suffix = stringSections[len(stringSections)-1]
	}

	for i := 0; i < len(stringSections)-1; i++ {
		name += stringSections[i]

		// adds back any stripped '.'s apart from the last one
		if i != len(stringSections)-2 {
			name += "."
		}
	}

	return
}

func ConstructPath(preffix, directory, fileName string) (s string) {
	if !reflect.DeepEqual(preffix, "") {
		s += preffix + "/"
	}

	s += directory

	if !reflect.DeepEqual(fileName, "") {
		s += "/" + fileName
	}
	return s
}

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
