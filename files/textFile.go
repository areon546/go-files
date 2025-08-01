package files

import (
	"errors"
	"os"
	"strings"

	"github.com/areon546/go-helpers/helpers"
)

// Contract:
//   - you can use me to write to files.
type TextFile struct {
	File
	textBuffer []string
	lines      int
}

func NewTextFile(filePath string) *TextFile {
	return &TextFile{File: *NewFile(filePath)}
}

func (f *TextFile) Contents() []string {
	return f.textBuffer
}

func (f *TextFile) ReadFile() []string {
	if !f.hasBeenRead {
		data, err := os.ReadFile(f.Name()) // For read access.
		handle(err)

		oneLine := strings.ReplaceAll(string(data), "\r", "")
		f.textBuffer = strings.Split(oneLine, "\n")
		f.lines = len(f.contentBuffer)
	}
	return f.textBuffer
}

func (f *TextFile) ReadLine(lineNum int) (output string, err error) {
	lineNum -= 1 // converted to index notation

	if f.IsEmpty() {
		f.ReadFile()
	}

	if lineNum > f.lines {
		return "", errors.New("Index out of bounds for File length")
	}

	output = string(f.textBuffer[lineNum])
	// print(output)

	return
}

func (t *TextFile) WriteLine(s string, i int, newline bool) {
	for i >= len(t.textBuffer) {
		t.textBuffer = append(t.textBuffer, "")
	}

	if newline {
		s += "\n"
	}

	t.textBuffer[i] = s
	t.File.Append(helpers.StringToBytes(s))
}

func (t *TextFile) WriteBuffer() {
	// convert bytes buffer to bytes
	textAsOneLine := strings.Join(t.textBuffer, "\n")

	bytes := helpers.StringToBytes(textAsOneLine)

	// write text buffer
	t.Write(bytes)
}

// ~~~ Append
func (t *TextFile) AppendLastLine(s string) {
	lastLine := len(t.contentBuffer) - 1

	if t.IsEmpty() {
		lastLine = 0
	}

	t.WriteLine(s, lastLine, true)
}

func (t *TextFile) AppendLines(arr []string, newline bool) {
	for _, v := range arr {
		t.Append(v, newline)
	}
}

func (f *TextFile) Append(s string, newline bool) {
	f.WriteLine(s, len(f.contentBuffer), newline)
}

func (f *TextFile) AppendNewLine(s string) {
	f.Append(s, true)
}

func (t *TextFile) AppendEmptyLine() {
	t.Append("", true)
}
