package files

import (
	"errors"
	"strings"

	"github.com/areon546/go-helpers/helpers"
)

// Contract:
//   - you can use me to write to files.
type TextFile struct {
	File       // May consider making this an attribute, rather than embed it like this
	textBuffer []string
	lines      int
}

func NewTextFile(filePath string) *TextFile {
	return &TextFile{File: *NewFile(filePath)}
}

// Reading and Writing from and to the TextFile.

func (t *TextFile) ReadContents() ([]string, error) {
	bytes, err := t.File.ReadContents()
	if err != nil {
		return []string{}, err
	}

	t.textBuffer = t.deserialise(bytes)
	t.lines = len(t.contentBuffer)

	return t.textBuffer, nil
}

// Convert from byte format (lower level struct) to human readable strings (struct )
func (f *TextFile) deserialise(bytes []byte) []string {
	oneLine := strings.ReplaceAll(string(bytes), "\r", "")
	// NOTE: This was done because Windows uses CR LF to denote line ends, while linux only uses LF
	// the only time this should break, should be in a system that uses CR, \r, to denote new lines,
	// however so far as I understand, that is only reserved to really old OS'

	strings := strings.Split(oneLine, "\n")
	return strings
}

func (t *TextFile) ReadLine(lineNum int) (output string, err error) {
	lineNum -= 1 // converted to index notation

	if t.IsEmpty() {
		t.ReadContents()
	}

	if t.IsEmpty() {
		// If still empty, TODO: return error
	}

	if lineNum > t.lines {
		return "", errors.New("Index out of bounds for File length")
	}

	output = string(t.textBuffer[lineNum])
	// print(output)

	return
}

func (t *TextFile) WriteContents() {
	bytes := t.serialise(t.textBuffer)

	t.File.Append(bytes)
	t.File.WriteContents()
}

// Converts from the TextFile struct data format ([]string), to the data format down the chain,
func (t *TextFile) serialise(lines []string) []byte {
	textAsOneLine := strings.Join(t.textBuffer, "\n")
	return helpers.StringToBytes(textAsOneLine)
}

// Misc

func (f *TextFile) Contents() []string {
	return f.textBuffer
}

// ~~~ Append

func (t *TextFile) SetLine(s string, i int, newline bool) {
	for i >= len(t.textBuffer) {
		t.textBuffer = append(t.textBuffer, s)
	}

	if newline {
		s += "\n"
	}

	t.textBuffer[i] = s
}

func (f *TextFile) Append(s string, newline bool) {
	f.SetLine(s, len(f.textBuffer), newline)
}

func (t *TextFile) AppendLastLine(s string) {
	lastLine := len(t.textBuffer) - 1

	if t.IsEmpty() {
		lastLine = 0
	}

	t.SetLine(s, lastLine, true)
}

func (t *TextFile) AppendLines(arr []string, newline bool) {
	for _, v := range arr {
		t.Append(v, newline)
	}
}

func (f *TextFile) AppendNewLine(s string) {
	f.Append(s, true)
}

func (t *TextFile) AppendEmptyLine() {
	t.Append("", true)
}
