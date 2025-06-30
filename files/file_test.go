package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

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

			print("TEST TEST TEST", file.filename)

			helpers.AssertEquals(t, tC.expected, name)
		})
	}
}
