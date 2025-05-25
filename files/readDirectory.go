package files

import (
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
