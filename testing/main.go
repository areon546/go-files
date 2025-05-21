package main

import (
	"os"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

func main() {
	// lookAtDirectory()

	files.OpenFile("testing/files/test.txt")

	// f.Close()
}

func lookAtDirectory() {
	fold, _ := os.ReadDir("/etc")
	helpers.Print(fold)

	for i := 0; i < len(fold); i++ {
		helpers.Print(fold[i])
	}
}
