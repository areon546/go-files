package files

import (
	"bytes"
	"fmt"
	"os"
	"syscall"
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

var (
	realName string = "./files/test.txt"
	realFile File

	duplicateName string = "./files/test2.txt"
	duplicateReal File

	fakeName string = "./files/faketest.txt"
	fakeFile File

	emtpyFile File
)

func init() {
	// Set up values for tests to make
	print("Setting up file_test")
	defer print("Finished setting up file_test\n")

	realFile = *NewFile(realName)
	realFile.ClearFile()
	realFile.Append([]byte("[dragons are cool]"))

	duplicateReal = *NewFile(duplicateName)
	duplicateReal.ClearFile()

	fakeFile = *NewFile(fakeName)

	emtpyFile = *EmptyFile()
}

/* ~~~ Creating File objects */

// NewFile lets you create digital file objects
func TestNewFile(t *testing.T) {
	// test that the path given to the file is appropriate
	// test that files start empty unless reading

	t.Run("New File", func(t *testing.T) {
		nFile := NewFile("")
		err := nFile.WriteContents()

		helpers.AssertError(t, err, ErrNoFileOrDirectory)
		helpers.AssertError(t, err, syscall.ENOENT) // ErrNoFileOrDirectory is a wrapper for syscall ENOENT
	})

	t.Run("Non-existant File", func(t *testing.T) {
		f := NewFile("files/faketest.txt")
		want := []byte{}
		get, err := f.ReadContents()

		// NOTE: Since asserting the actual error is difficult, and I have tried but been able to find the relevant error for
		helpers.AssertError(t, err, ErrNoFileOrDirectory)
		helpers.AssertEqualsBytes(t, want, get)

		fmt.Println("ASDASD", bytes.Equal(want, get))
	})
}

// OpenFile lets you load files from the filesys to memory
func TestOpenFile(t *testing.T) {
	t.Run("Open Existing File", func(t *testing.T) {
		_, err := OpenFile("files/test.txt")

		// when we open a file, id expect the buffer to contain the file's contents, and reading from the file to have no affect
		//

		// assert there is an error
		helpers.AssertNoError(t, err)
	})

	t.Run("Non-existant File", func(t *testing.T) {
		_, err := OpenFile("files/faketest.txt")

		helpers.AssertError(t, err, os.ErrNotExist)
	})
}

// EmptyFile is a 'nil' file for comparisons, used internally, to ensure
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

// ReadContents reads the already present file content
func TestFileReadContents(t *testing.T) {
}

// WriteContents writes the contents of the buffer onto
func TestFileWriteContents(t *testing.T) {
}

// func TestClearFile(t *testing.T) {
// 	// test file with data written to the file, data will be overwritten
// 	// t.Run("File Emptied", func(t *testing.T) {
// 	// })
// 	// test file with data in buffer that buffer will not be overwritten
// 	t.Run("File Emptied but buffer not Emptied", func(t *testing.T) {})
// }
//
// func TestWrite(t *testing.T) {
// 	t.Run("", func(t *testing.T) {})
// }
//
// func TestRead(t *testing.T) {
// 	t.Run("", func(t *testing.T) {})
// }
//
// /* ~~~ Misc */
//
// // IsEmpty
// func TestIsEmpty(t *testing.T) {
// 	t.Run("", func(t *testing.T) {})
//
// 	t.Run("", func(t *testing.T) {})
// }
//
// // Contents
// func TestConentsFile(t *testing.T) {
// 	t.Run("", func(t *testing.T) {})
// }
//
// // String
// func TestString(t *testing.T) {
// 	t.Run("Test String returns path to file object", func(t *testing.T) {
// 	})
// }

func TestRename(t *testing.T) {
	f := realFile
	o := realFile
	helpers.AssertEquals(t, "./files/test.txt", realFile.FullName())

	new := "files/test2.txt"
	path, fn := SplitFilePath(new)
	f.Rename(path, fn)

	helpers.AssertEquals(t, "./files/test2.txt", f.FullName()) // Test the internal file path has changed
	helpers.AssertEquals(t, "./files/test.txt", o.FullName())

	// Test it will write and not overwrite the original
	err := f.WriteContents()
	helpers.AssertNoError(t, err)

	f2, err := OpenFile(new)
	helpers.AssertNoError(t, err)
	helpers.AssertEquals(t, string(o.Contents()), string(f2.Contents()))
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
			expected: "file.txt",
		},
		{
			desc:     "without /",
			path:     "file.txt",
			expected: "file.txt",
		},
		{
			desc:     "empty",
			path:     "",
			expected: "",
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
	testCases := []struct {
		desc     string
		path     string
		expected string
	}{
		{
			desc:     "with /",
			path:     "./file.txt",
			expected: "./",
		},
		{
			desc:     "without /",
			path:     "file.txt",
			expected: "./",
		},
		{
			desc:     "empty",
			path:     "",
			expected: "./",
		},
		{
			desc:     "present, as test ",
			path:     "/etc/passwd",
			expected: "/etc/",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// open file
			file := NewFile(tC.path)
			name := file.Path()

			helpers.AssertEquals(t, tC.expected, name)
		})
	}
}

// Path
func TestFullName(t *testing.T) {
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
		{
			desc:     "empty",
			path:     "",
			expected: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// open file
			file := NewFile(tC.path)
			name := file.FullName()

			helpers.AssertEquals(t, tC.expected, name)
		})
	}
}
