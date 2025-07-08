package files

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

var (
	errFiles        = errors.New("files: ")
	errNotDirectory = errors.New("specified path not a directory, must end with '/': ")

	ErrNotDir = fmt.Errorf("%w%w", errFiles, errNotDirectory)
)

// This file contains all of the methods relating to directory management and checking.

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
		exists = true && !info.IsDir()
	}

	return
}

// Checks if a specific directory exists.
func DirExists(path string) (exists bool, info os.FileInfo) {
	info, err := os.Stat(path)

	if errors.Is(err, fs.ErrNotExist) {
		exists = false
	} else {
		exists = true && info.IsDir()
	}

	return
}

// Creates directories at the specified path.
// Returns an error if there is an issue creating the directories, or if the path is not a directory path.
func MakeDirectory(path string) error {
	if PathIsDir(path) {
		return os.MkdirAll(path, os.ModePerm)
	} else {
		return errors.Join(ErrNotDir, errors.New(path))
	}
}

// Checks if the path specified is that of a directory.
func PathIsDir(path string) bool {
	return path[len(path)-1] == '/'
}

func CleanUpDirs(dirs []string) []string {
	newDirs := make([]string, 0)

	for index, dir := range dirs {
		print(index, dir)

		newDirs := 
	}

	return nil
}
