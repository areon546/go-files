package files

import (
	"io/fs"
	"log"
	"os"

	"github.com/areon546/go-helpers/helpers"
)

// This function
func ReadDirectory(dirPath string) (entries []fs.DirEntry) {
	helpers.Printf("Reading directory %s", dirPath)

	// read contents of a directory
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	return
}
