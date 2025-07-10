package files

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"reflect"
	"regexp"

	"github.com/areon546/go-helpers/helpers"
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

// Returns processes a dirs array, such as one returned by files.SplitDirectories.
// Main consequences:
//   - removes any "." directories due to redundancy
//
// Should this return an error if relative imports go up more times than exist within the function?
func CleanUpDirs(dirs []string) []string {
	newDirs := make([]string, 0)

	for _, dir := range dirs {
		validDirectory := ValidDirectoryName(dir)

		debugPrint(helpers.Format("CleanUpDirs processing: %t '%s'", validDirectory, dir))

		if validDirectory {
			newDirs = append(newDirs, dir)
		}
	}

	return newDirs
}

func ValidDirectoryName(dir string) (isValid bool) {
	containsOnlySpacesRegex := "\\ [\\ ]*" // \ [\ ]*

	dotDirectory := reflect.DeepEqual(dir, ".")

	validCharacters := validPosixName(dir) // regex check for whether it contains only valid characters

	onlySpaces, err := regexp.MatchString(containsOnlySpacesRegex, dir) // regex check for whether it contains only spaces
	helpers.Handle(err)

	isValid = !dotDirectory && !onlySpaces && validCharacters

	debugPrint("ValidDirectoryName output", isValid, ":", !dotDirectory, !onlySpaces, validCharacters)

	return
}

func ValidFileName(file string) (isValid bool) {
	return
}

func validPosixName(name string) (isValid bool) {
	validCharRegex := "[A-Za-z0-9\\.\\_\\-]+" // [A-Za-z0-9\.\_\-]+
	// Follows the POSIX Portable Character
	// https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap03.html#tag_03_282

	validCharacters, err := regexp.MatchString(validCharRegex, name) // regex check for whether it contains only valid characters - gonna make it
	if err != nil {
		return false
	}

	isValid = validCharacters

	return isValid
}
