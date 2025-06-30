package files

import (
	"errors"
	"io/fs"
	"log"
	"os"
)

// This function
func ReadDirectory(dirPath string) (entries []fs.DirEntry) {
	debugPrint("Reading directory ", dirPath)

	// read contents of a directory
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Checks if a specific file exists within the file system or not.
// If it does exist, it returns true, and the os.FileInfo entry that can be gained by os.Stat.
// If it does not exist, it returns false, nil.
func FileExists(path string) (exists bool, info os.FileInfo) {
	info, err := os.Stat(path)

	if errors.Is(err, fs.ErrNotExist) {
		exists = false
	} else {
		exists = true
	}

	return
}

// TODO: make directory function?
