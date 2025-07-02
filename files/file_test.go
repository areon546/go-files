package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

var (
	realFile  File = *NewFile("test.txt")
	fakeFile  File = *NewFile("faketest.txt")
	emtpyFile File = *EmptyFile()
)

func init() {
	// Set up values for tests to make
}

/* ~~~ Creating File objects */

// NewFile
func TestNewFile(t *testing.T) {
	// t.Run("New File", func(t *testing.T) {
	// 	nFile := NewFile("")
	// 	nFile.Close()
	// })
	//
	//
	// t.Run("Non-existant File", func(t *testing.T) {
	// 	OpenFile("files/faketest.txt")
	// })
}

// OpenFile
func TestOpenFile(t *testing.T) {
	t.Run("New File", func(t *testing.T) {
		OpenFile("files/test.txt")

		// when we open a file, id expect the buffer to contain the file's contents, and reading from the file to have no affect
	})

	// t.Run("Non-existant File", func(t *testing.T) {
	// 	OpenFile("files/faketest.txt")
	// })
}

// EmptyFile
func TestEmptyFile(t *testing.T) {
	// representation of an empty file for comparisons and handling, and the like

	t.Run("Test empty file contains nothing inside of it", func(t *testing.T) {
		helpers.AssertEqualsObject(t, &File{}, EmptyFile())
	})
}

/* ~~~ Reading and Writing */

func TestClearFile(t *testing.T) {
	// test file with data written to the file, will be deleted
	// test file with data in buffer that buffer will not be deleted
	t.Run("", func(t *testing.T) {})
}

func TestWrite(t *testing.T) {
	t.Run("", func(t *testing.T) {})
}

func TestRead(t *testing.T) {
	t.Run("", func(t *testing.T) {})
}

/* ~~~ Misc */

// IsEmpty
func TestIsEmpty(t *testing.T) {
	t.Run("", func(t *testing.T) {})

	t.Run("", func(t *testing.T) {})
}

// Contents
func TestConentsFile(t *testing.T) {
	t.Run("", func(t *testing.T) {})
}

// String
func TestString(t *testing.T) {
	t.Run("Test String returns path to file object", func(t *testing.T) {
	})
}

// Name
func TestName(t *testing.T) {
	testCases := []struct {
		desc     string
		path     string
		expected string
	}{
		{
			desc:     "with /",
			path:     "./file.txt",
			expected: "./file.txt",
		},
		{
			desc:     "without /",
			path:     "file.txt",
			expected: "./file.txt",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// open file
			file := NewFile(tC.path)
			name := file.Name()

			helpers.AssertEquals(t, tC.expected, name)
		})
	}
}

// Path
func TestPath(t *testing.T) {
}
