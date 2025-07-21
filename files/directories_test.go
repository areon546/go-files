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
		{
			desc:     "Empty Directories",
			path:     "",
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			helpers.AssertEqualsBool(t, PathIsDir(tC.path), tC.expected)
		})
	}
}

func TestCleanUpDirs(t *testing.T) {
	t.Run("", func(t *testing.T) {})

	testCases := []struct {
		desc         string
		inputDirs    []string
		expectedDirs []string
	}{
		{
			desc:         "Test Stripping directories of '.'",
			inputDirs:    []string{"asdasd", ".", ".", "asdasdasd"},
			expectedDirs: []string{"asdasd", "asdasdasd"},
		},
		{
			desc:         "Test Stripping directories of (SPACE)",
			inputDirs:    []string{"aa", " ", "asd", "aaas"},
			expectedDirs: []string{"aa", "asd", "aaas"},
		},
		// {
		// 	desc:         "",
		// 	inputDirs:    []string{},
		// 	expectedDirs: []string{},
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cleanedUpDirs := CleanUpDirs(tC.inputDirs)

			helpers.AssertEqualsObject(t, tC.expectedDirs, cleanedUpDirs)
		})
	}
}

func TestValidDirectoryName(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc:     "Forbid (SPACE) when alone (One)",
			input:    " ",
			expected: false,
		},
		{
			desc:     "Forbid (SPACE) when alone (Many)",
			input:    "       ",
			expected: false,
		},
		{
			desc:     "Forbid starting with (SPACE)",
			input:    " asd",
			expected: false,
		},
		{
			desc:     "Forbid ending with (SPACE)",
			input:    "asd ",
			expected: false,
		},
		{
			desc:     "Forbid . as only character",
			input:    ".",
			expected: false,
		},
		{
			desc:     "Allow '.' when not alone",
			input:    "directories.pretending.to.be.files.is.annoying",
			expected: true,
		},
		{
			desc:     "Allow '-'",
			input:    "kebabs-are-tasty",
			expected: true,
		},
		{
			desc:     "Allow '_'",
			input:    "words_are_powerful",
			expected: true,
		},
		// {
		// 	desc:     "",
		// 	input:    "",
		// 	expected: false,
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			isValid := ValidDirectoryName(tC.input)

			// print(isValid, tC)
			helpers.AssertEqualsBool(t, tC.expected, isValid)
		})
	}
}
