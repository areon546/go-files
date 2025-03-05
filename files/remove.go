package files

import "os"

// I want to remove all the files within a given directory.
//  but keep the directory

// RemoveAllWithinDirectory removes all the files in the specified directory.
func RemoveAllWithinDirectory(path string) {
	os.RemoveAll(path)          // NOTE: Removes the directory in the path specified.
	os.Mkdir(path, os.ModePerm) // Adds back the removed directory.
}
