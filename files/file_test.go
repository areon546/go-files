package files

import (
	"os"
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

	realFile.Append([]byte("[dragons are cool]"))
}

/* ~~~ Creating File objects */

// NewFile
func TestNewFile(t *testing.T) {
	// test that the path given to the file is appropriate
	// test that files start empty unless reading

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
	t.Run("Open Existing File", func(t *testing.T) {
		_, err := OpenFile("files/test.txt")

		// when we open a file, id expect the buffer to contain the file's contents, and reading from the file to have no affect
		//

		// assert there is an error
		helpers.AssertError(t, err, nil)
	})

	t.Run("Non-existant File", func(t *testing.T) {
		_, err := OpenFile("files/faketest.txt")

		helpers.AssertError(t, err, os.ErrNotExist)
	})
}

// EmptyFile
func TestEmptyFile(t *testing.T) {
	// representation of an empty file for comparisons and handling, and the like
	empty := EmptyFile()

	t.Run("Empty file contains nothing inside of it", func(t *testing.T) {
		// assert that it is empty
		helpers.AssertEqualsObject(t, &File{file: file{}}, empty)

		// assert name is empty
		// helpers.AssertEquals(t, "", empty.Path()) // tests fail, actual is ./
		// helpers.AssertEquals(t, "", empty.Name()) // tests fail, actual is ./. (path + . )    - do i want this?

		// assert content buffer is empty on initialisation
		helpers.AssertEquals(t, "", string(empty.Contents()))

		// assert unable to read - literally forbiddden because path is missing, ditto for writing
		// I would like if it gave an error, however, do i technically want an error?
		//
		// Bad tests
		b := make([]byte, 50)
		emptyBytes := make([]byte, 50)
		_, err := empty.Read(b)

		helpers.AssertError(t, err, ErrPathEmpty)
		helpers.AssertEquals(t, string(emptyBytes), string(b))
		helpers.AssertEqualsInt(t, len(emptyBytes), len(b))

		// assert unable to write

		// * empty name
		// * empty content buffer
		// * unable to read
		// * unable to write
	})
}

/* ~~~ Reading and Writing */

func TestClearFile(t *testing.T) {
	// test file with data written to the file, data will be overwritten
	// t.Run("File Emptied", func(t *testing.T) {
	// })
	// test file with data in buffer that buffer will not be overwritten
	t.Run("File Emptied but buffer not Emptied", func(t *testing.T) {})
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
