package files_test

import (
	"testing"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

func TestIsEmpty(t *testing.T) {

	t.Run("", func(t *testing.T) {})

	t.Run("", func(t *testing.T) {})

}

func TestName(t *testing.T) {

	t.Run("", func(t *testing.T) {})

}

func TestConentsFile(t *testing.T) {

	t.Run("", func(t *testing.T) {})

}

func TestClearFile(t *testing.T) {

	t.Run("", func(t *testing.T) {})

}

func TestWrite(t *testing.T) {

	t.Run("", func(t *testing.T) {})

}

func TestRead(t *testing.T) {

	t.Run("", func(t *testing.T) {})

}

func TestString(t *testing.T) {

	t.Run("Test String returns path to file object", func(t *testing.T) {

	})

}

func TestEmptyFile(t *testing.T) {
	// representation of an empty file for comparisons and handling, and the like

	t.Run("Test empty file contains nothing inside of it", func(t *testing.T) {
		helpers.AssertEqualsObject(t, &files.File{}, files.EmptyFile())
	})
}

func TestOpenFile(t *testing.T) {

	t.Run("New File", func(t *testing.T) {
		files.OpenFile("files/test.txt")

		// when we open a file, id expect the buffer to contain the file's contents, and reading from the file to have no affect
	})

	t.Run("Non-existant File", func(t *testing.T) {
		files.OpenFile("files/faketest.txt")

	})

	t.Run("New File", func(t *testing.T) {
		nFile := files.NewFile("")
		nFile.Close()
	})
}
