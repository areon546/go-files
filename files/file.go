package files

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"syscall"

	"github.com/areon546/go-helpers/helpers"
)

// ~~~~~~~~~~~~~~~~ File

// The File struct is supposed to handle the io.ReadWriter interface.
// It treats files as binary objects, writing to them with bytes.
//
// There are two ways to iteract with this file system
//  1. Calling file.Write(byteArr) will write directly the contents of the byte array into the file.
//  2. Calling Append will buffer information into the content buffer, allowing the struct to act as a BufferedWriter in Java
//     Note: When appending to the buffer, at the end you have to
//
// Promises of the File struct
// - it will fulfill the io.ReadWriteCloser interface
type File struct {
	file
}

type file struct {
	path     string // path attribute has to be a directory, ie ending with a `/`
	filename string

	contentBuffer    []byte
	compiledFilePath string

	bytesRead int
}

/*
File Creation:
- NewFile - creates a fake file that needs to be written or closed to actually appear in the file system
- OpenFile - reads from the OS, looking for a pre existing file with specified name, will buffer the file's contents into memory
- EmptyFile - creates an uneditable file

*/

// Main Mehtods:
/*
Read - will read contents into specified array
Write - will overwrite contents with specified array

Append - appends bytes to buffer
Close - writes buffer

*/

/* ~~~~~~~~ Creating File objects ~~~~~~~~ */

// Loads file from memory, loading any contents into the file created.
// The intended way to create files if they exist within memory.
func OpenFile(path string) (f *File, err error) {
	f = NewFile(path)
	_, err = f.ReadContents()
	return f, err
}

// Creates a new file at a specified directory.
// Promises of a 'NewFile':
// * will be created regardless
// * the new file will have empty contents, inte
func NewFile(filePath string) *File {
	debugPrint("Creating new file at path: ", filePath)
	path, fn := SplitFilePath(filePath)

	f := file{path: path, filename: fn}
	f.contentBuffer = []byte{}

	debugPrint("New File created", filePath, &f, "filename:", fn, "filename:", f.filename)
	return &File{f}
}

// Returns a basic, empty file for use in comparisons when a method has to return a file but has an inssue. Equivalent of `nil`.
// Promises of an Empty File:
// * empty name
// * empty content buffer
// * unable to read
// * unable to write
func EmptyFile() *File {
	return &File{file: file{}}
}

/* ~~~~~ Writing and Reading ~~~~~~~ */

// The 0664 is the permissions the file is written to with, however you can encode some additional stuff with it that isn't currently considered.

// Writes to the specified file.
func writeToFile(filename string, bytes []byte) error {
	return os.WriteFile(filename, bytes, 0o644)
}

// Appends to the specified file.
func appendToFile(filename string, bytes []byte) (err error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}

	_, err = f.Write(bytes)
	return err
}

// copied from io.go
//
// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
//
// Implementations must not retain p.

// Fulfills io.Writer interface.
// Writes to the file, appending the given bytes.
func (f *File) Write(p []byte) (n int, err error) {
	err = appendToFile(f.FullName(), p) // NOTE: this hass to append to fulfill the html template system, presumably stream means appending in this case
	if err == nil {                     // tell the user that the file has been written to successfully, only if no error occurs
		n = len(p) // simplistic answer
	}
	return n, err
}

// copied from io.go
//
// Read reads up to len(p) bytes into p. It returns the number of bytes
// read (0 <= n <= len(p)) and any error encountered. Even if Read
// returns n < len(p), it may use all of p as scratch space during the call.
// If some data is available but not len(p) bytes, Read conventionally
// returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after
// successfully reading n > 0 bytes, it returns the number of
// bytes read. It may return the (non-nil) error from the same call
// or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next Read should
// return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// If len(p) == 0, Read should always return n == 0. It may return a
// non-nil error if some error condition is known, such as EOF.
//
// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.
//
// Implementations must not retain p.

// Fulfills io.Reader interface.
// Reads the file, and returns the number of bytes read.
func (f *File) Read(p []byte) (n int, err error) {
	pathEmpty := reflect.DeepEqual(f.path, "")
	if pathEmpty {
		return 0, ErrPathEmpty
	}

	l := len(f.contentBuffer)
	i := 0

	for ; i < l; i++ {
		v := f.contentBuffer[i]
		if i < len(p) {
			p[i] = v
			n++
		} else if i == len(p) {
			// err
			err = io.EOF
			break
		}
	}

	if i == l {
		err = io.EOF
	}

	f.bytesRead += n
	// helpers.Print(l, len(p), f.Name(), "bytes read", f.bytesRead)
	// time.Sleep(time.Millisecond * 400)
	return n, err
}

// Read the contents of the file into the file buffer.
// Allows a user to create a file using `NewFile` and then later load the contents into the buffer as desired.
// Returns:
// - os.ErrNotExist if the file object does not exist in the file system.
func (f *File) ReadContents() ([]byte, error) {
	contents, err := f.deserialise()
	f.Append(contents)

	// Error check for errors similar to syscall.ENOENT "no such file or directory"
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		fmt.Println(f.Name(), err)

		err = newErrNoFileOrDirectory(f.FullName())
		fmt.Println("ASDASDASD")
	}
	return contents, err
}

// Load data from file to struct
func (f *File) deserialise() ([]byte, error) {
	bytes, err := os.ReadFile(f.FullName())

	return bytes, err
}

// WriteContents writes the content buffer to the file in the file system.
func (f *File) WriteContents() error {
	return f.serialise(f.Contents())
}

// Load data in struct buffer to file
func (f *File) serialise(bytes []byte) error {
	err := f.ClearFile()
	if err != nil {

		isNoFileOrDir := errors.Is(err, syscall.ENOENT)
		if isNoFileOrDir { // EG: if the name of the file is "", this returns 'no such file or directory'
			return ErrNoFileOrDirectory
		}
		return helpers.WrapError("files: aborted due to error %s", err)
	}

	_, err = f.Write(bytes)
	return err
}

// Appends to the end of the buffer.
func (f *File) Append(bytes []byte) {
	// adds bytes array given to the end of the buffer
	f.contentBuffer = append(f.contentBuffer, bytes...)
}

// Resets the actual file's contents. Helpful if writing to a pre-existing file and you don't care about the original content.
func (f *File) ClearFile() error {
	return writeToFile(f.FullName(), make([]byte, 0))
}

/* Misc methods */

// Renames the filename of the file object.
func (f *File) Rename(path, name string) {
	f.path = path
	f.filename = name
}

// Returns if the file buffer is empty.
func (f *File) IsEmpty() bool {
	return len(f.contentBuffer) == 0
}

// Returns the contents stored in the file buffer.
func (f *File) Contents() []byte {
	return f.contentBuffer
}

// Returns Important information about the file to be able to gather more information when printing.
func (f *File) String() string {
	size := len(f.Contents())

	return "Name:" + f.Name() + "Size:" + helpers.IntegerToString(size)
}

// Returns the file name.
func (f *File) Name() string {
	return f.filename
}

// Returns the directory of the specified file.
// Promise: Must end with a '/'
func (f *File) Path() string {
	path := f.path

	emptyPath := reflect.DeepEqual(path, "")
	pathPreffix := strings.HasPrefix(path, "./") || strings.HasPrefix(path, "/")
	if emptyPath || !pathPreffix {
		path = "./" + path
	}

	// check if path ends with a "/"
	if !PathIsDir(path) {
		path += "/"
	}

	f.compiledFilePath = path

	return f.compiledFilePath
}

func (f File) FullName() string {
	if f.filename == "" {
		return ""
	}

	return f.Path() + f.filename
}
