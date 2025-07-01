package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func TestReadDirectory(t *testing.T) {
}

func TestDirExists(t *testing.T) {
}

func TestMakeDirectory(t *testing.T) {
	testCases := []struct {
		desc     string
		path     string
		expected error
	}{
		{
			desc:     "Valid directory, ending with /",
			path:     "asd/",
			expected: nil,
		},
		{
			desc:     "Invalid directory, ending without /",
			path:     "asd",
			expected: ErrNotDir,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			helpers.AssertError(t, MakeDirectory(tC.path), tC.expected)
		})
	}
}

func TestPathIsDir(t *testing.T) {
	testCases := []struct {
		desc     string
		path     string
		expected bool
	}{
		{
			desc:     "Valid directory, ending with /",
			path:     "asd/",
			expected: true,
		},
		{
			desc:     "Invalid directory, ending without /",
			path:     "asd",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			helpers.AssertEqualsBool(t, PathIsDir(tC.path), tC.expected)
		})
	}
}
