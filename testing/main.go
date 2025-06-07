package main

import (
	"github.com/areon546/go-files/files"
)

func main() {
	// lookAtDirectory()

	f := files.OpenFile("testing/files/test.txt")
	_, err := f.Write([]byte("hi test"))

	print(err)
	// f.Close()
}
