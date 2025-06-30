package files

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

// This will join together the path, filename, and a specified file type.
func ConstructFilePath(path, filename, filetype string) (filePath string) {
	filePath = path + "/" + filename + filetype

	debugPrint("ConstructFilePath: params: ", path, filename, filetype)
	debugPrint("ConstructFilePath: return: ", filePath)

	return
}

// This will split up a file path into it's constituting directories, and filename.
func SplitDirectories(filePath string) (dirs []string, filename string) {
	stringSections := strings.Split(filePath, "/")
	length := len(stringSections)
	debugPrint("SplitDirectories: file path", filePath, ", split up: ", stringSections)

	// Check if filePath is a directory or a file
	filePresent := false
	info, err := os.Stat(filePath)

	// os.Stat returns either:
	// nil Err

	if err != nil {
		debugPrint("SplitDirectories: os.Stat Error", err)

		if errors.Is(err, fs.ErrNotExist) {
			// If the filePath object doesn't exist, we will assume that if the last stringSection of the filePath has a dot, it will be a file.

			if strings.Contains(stringSections[length-1], ".") {
				filePresent = true
			} // TODO: causes bug where if the directory somehow contains a dot in it, it is possible (sometimes) that an error will be formed.
		} else {
			handle(err)
		}
	} else if !info.IsDir() {
		filePresent = true
	}

	if filePresent {
		dirs = stringSections[0 : length-1]
		filename = stringSections[length-1]
	} else {
		dirs = stringSections[0 : length-0]
		filename = ""
	}

	// When the filePath is specifically that of a directory, and ends with a "/", dirs ends with an empty string.
	// The empty string can be removed.
	lenDir := len(dirs)
	if dirs[lenDir-1] == "" {
		dirs = dirs[0 : lenDir-1]
	}

	debugPrint("SplitDirectories return values: ", dirs, filename)
	return
}

// This will isolate a directory path, a file name, and the file type from a specified file path.
// If given a single line, it will
func SplitFilePath(filePath string) (path, name, filetype string) {
	dirs, filename := SplitDirectories(filePath)
	length := len(dirs)
	debugPrint("SplitFilePath dirs:", dirs, "filename", filename)

	if length > 0 {
		// This means that directories have been entered that we have to process and add onto the path string

		for i := range length {
			dir := dirs[i]

			debugPrint("SplitFilePath: appending / to each dir: ", i, dir)

			path += dir + "/"
		}
	}

	name, filetype = splitFileName(filename)
	debugPrint("SplitFilePath: ", "fn", filename, "name", name, "suff", filetype)

	debugPrint("Return values: ", path, name, filetype)
	return path, name, filetype
}

// Name is the whole name up to the very very last .xxx at the end of a filename.
// EG asd.jar.JAR.jar.txt will consider the filename as asd.jar.JAR.jar and the type as txt
func splitFileName(filename string) (name, filetype string) {
	// split it up by .'s
	startOfFileType := strings.LastIndex(filename, ".")

	debugPrint("splitFileName inputs:", filename, startOfFileType)

	if startOfFileType == -1 {
		// IE no actual . found in `filename`.
		return filename, ""
	}

	name = filename[0:startOfFileType]
	filetype = filename[startOfFileType+1:]
	debugPrint("splitFileName outputs:", name, filetype)
	return
}
