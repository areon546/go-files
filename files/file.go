package files

import (
	"io"
	"log"
	"os"

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
type File struct {
	path     string
	filename string
	suffix   string

	contentBuffer []byte
	lines         int
	linesRead     int
	hasBeenRead   bool

	bytesRead int
}

// Main Mehtods:
/*
Read - will read contents into specified array
Write - will overwrite contents with specified array

Append - appends bytes to buffer
Close - writes buffer



*/

func NewFile(filePath string) *File {
	path, fn, suff := SplitFilePath(filePath)

	f := &File{path: path, filename: fn, suffix: suff}

	f.hasBeenRead = false
	f.linesRead = 0

	return f
}

func (f *File) IsEmpty() bool {
	return len(f.contentBuffer) == 0
}

func (f *File) Name() string {
	return format("%s/%s.%s", f.path, f.filename, f.suffix)
}

func (f *File) Contents() []byte {
	return f.contentBuffer
}

// The 0664 is the permissions the file is written to with, however you can encode some additional stuff with it that isn't currently considered.
func writeToFile(filename string, bytes []byte) error {
	return os.WriteFile(filename, bytes, 0664)
}

func (f *File) ClearFile() {
	err := writeToFile(f.Name(), make([]byte, 0))
	log.Fatal(err)

}

// Writes the content buffer
func (f *File) Close() {
	f.Write(f.Contents())
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
func (f *File) Write(p []byte) (n int, err error) {
	err = writeToFile(f.Name(), p)
	return
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
func (f File) Read(p []byte) (n int, err error) {
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
	return
}

func (f *File) String() string {
	return f.Name()
}

func (f *File) Append(bytes []byte) {
	// adds bytes array given to the end of the buffer
	f.contentBuffer = append(f.contentBuffer, bytes...)
}

// Returns a basic, empty file
func EmptyFile() *File {
	return &File{}
}

// Loads file from memory, loading any contents into the file created
func OpenFile(path string) (f *File) { // TODO make the File struct use byte slice rather than string slice

	contents, err := os.ReadFile(path)
	helpers.Handle(err)

	f = NewFile(path)

	f.Append(contents)

	return
}
