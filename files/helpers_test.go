package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func TestSplitFileName(t *testing.T) {
	// when splitting the filename, the cases we care about are:

	var testCases = []struct {
		path     string
		filename string
		suffix   string
	}{
		{"./test.txt", "./test", "txt"},
		{"./test.txt.txt", "./test.txt", "txt"},
		{"../test.txt", "../test", "txt"},
		{"/test.txt", "/test", "txt"},
		{"/test", "/test", ""},
	}

	for _, testCase := range testCases {
		path := testCase.path

		filename := testCase.filename
		suffix := testCase.suffix

		name := helpers.Format("Test case: %s", path)

		t.Run(name, func(t *testing.T) {
			fn, sf := splitFileName(path)

			helpers.AssertEquals(t, filename, fn)
			helpers.AssertEquals(t, suffix, sf)
		})
	}
}
