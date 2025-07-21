package files

import (
	"reflect"
	"strings"
)

// This will join together the path, filename, and a specified file type.
func ConstructFilePath(path, filename, filetype string) (filePath string) {
	if !PathIsDir(path) { // appends a `/` to the path variable
		if reflect.DeepEqual(path, "") {
			path = "."
		}
		path += "/"
	}

	filePath = path + filename + "." + filetype

	debugPrint("ConstructFilePath: params: ", path, filename, filetype)
	debugPrint("ConstructFilePath: return: ", filePath)

	return
}

// This will split up a file path into it's constituting directories, and filename.
func SplitDirectories(filePath string) (dirs []string, filename string) {
	debugPrint("SplitDirectories inputs:", filePath)

	lastDirIndex := strings.LastIndex(filePath, "/")

	if lastDirIndex == -1 {
		// No directories within file Path
		dirs = []string{}
		// NOTE: For some reason, the default initialisation of a []string in the function definition, is not equal to this.
		// And the tests fail if you don't reinitialise it.
		filename = filePath

		return
	}

	dirString := filePath[0:lastDirIndex]
	filename = filePath[lastDirIndex+1:]

	dirs = strings.Split(dirString, "/")

	debugPrint("SplitDirectories return values: ", dirs, filename)
	return
}

// This will isolate a directory path, a file name, and the file type from a specified file path.
// Rule: Directories end with `/`
// Rule: Files don't end with a `/`
// When processing directory or not, if the last one has a dot in it,
func SplitFilePath(filePath string) (path, name, filetype string) {
	dirs, filename := SplitDirectories(filePath)
	// NOTE: I could make this slightly more optimised by not turning the filepath into a string array and then joining it back up again.
	// However, there are some rules that I want to be sure are maintained, and there are more important things.
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

func JoinDirs(dirs []string) (dirPath string) {
	return strings.Join(dirs, "/") + "/"
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
