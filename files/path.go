package files

import (
	"strings"
)

// This will join together the path, filename, and a specified file type.
func ConstructFilePath(path, filename, filetype string) (filePath string) {
	return path + "/" + filename + filetype
}

// This will isolate a directory path, a file name, and the file type from a specified file path.
func SplitFilePath(filePath string) (path, name, filetype string) {
	stringSections := strings.Split(filePath, "/")
	length := len(stringSections)

	debugPrint("file name split up: ", stringSections)

	if length == 0 {
		// IE "" or "/" entered.
		return
	}

	if length == 1 {
		// This means that only a filename has been entered (eg TEST.MD).
		path = ""
	}

	// Split up filename

	if length > 1 {
		// This means that directories have been entered that we have to process and add onto the path string

		for i := range length - 1 {
			debugPrint(i, stringSections[i])

			path += stringSections[i]
		}
	}

	name, filetype = splitFileName(stringSections[length-1])

	debugPrint(name, filetype)

	return
}

// Name is the whole name up to the very very last .xxx at the end of a filename.
// EG asd.jar.JAR.jar.txt will consider the filename as asd.jar.JAR.jar and the type as txt
func splitFileName(filename string) (name, filetype string) {
	// split it up by .'s
	startOfFileType := strings.LastIndex(filename, ".")

	if startOfFileType == -1 {
		// IE no actual . found in `filename`.
		return filename, ""
	}

	name = filename[0:startOfFileType]
	filetype = filename[startOfFileType+1:]
	return
}
