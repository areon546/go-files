package files

import (
	"reflect"
	"strings"

	"github.com/areon546/go-helpers/helpers"
)

// helper functions

func handle(err error) {
	// helpers.Handle(err)
	print("files package handling error", err)
	helpers.Handle(err)
}

func print(a ...any) {
	helpers.Print(a...)
}

// unique ~~~~
func FilesEqual(a, b File) bool {
	return reflect.DeepEqual(a, b)
}

func HasSuffix(file, suffix string) bool {
	fileSuffix := "." + suffix
	return strings.HasSuffix(file, fileSuffix)
}

// Appends the filetype 'filetype' specified to the end of the 'file' string.
func AddFileType(file, filetype string) string {
	if HasSuffix(file, filetype) {
		return file
	}
	return file + "." + filetype
}

func trimFiletype(filename, filetype string) string {
	if filenameContains(filename, filetype) {
		return removeFiletype(filename, filetype)
	}
	return filename
}

func filenameContains(filename, filetype string) bool {
	filenameComponents := strings.Split(filename, ".")
	fileTypeInFileName := filenameComponents[len(filenameComponents)-1]

	return reflect.DeepEqual(fileTypeInFileName, filetype)
}

func removeFiletype(filename, filetype string) string {
	return strings.TrimSuffix(filename, ("." + filetype))
}

func returnErr(err error) error {
	if err != nil {
		errIsEmpty := err == errEmpty
		if !errIsEmpty {
			return err
		}
	}
	return nil
}
