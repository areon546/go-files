package files

import (
	"io/fs"
	"log"
	"os"
)

// This function
func ReadDirectory(dirPath string) (entries []fs.DirEntry) {
	// read contents of a directory
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	return
}
