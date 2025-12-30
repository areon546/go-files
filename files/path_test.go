package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

/*
Ok, what are the needs of a Path Object?

Do I want it to be a completely internal thing for my own processing or something for others to use aswell?
For the time being I don't actually see a need for other people to care about path objects so it really should be an internal object.

I want this to simplify my life when calling paths internally, the File class would just reference a path object, and then I can call methods on it like
file.path.getFileName
file.path.getPath
file.path.getFileType

Do I want it to be able to support absolute and relative paths?
This is a question of would the File class supprot reading from absolute paths.

Do I want it to support windows using forward slashes? Yes. No reason to be pedantic and only support UNIX,

*/

// func TestSplitPath(t *testing.T) {

// 	var testCases = []struct {
// 		message  string
// 		path     string
// 		expected []string
// 	}{
// 		// path splitting functionality
// 		{"Split path into array of directory names", "testing/files", []string{"testing", "files"}},

// 		// unix tests
// 		{"UNIX: Support absolute paths", "/testing/files", []string{"testing", "files"}},
// 		{"UNIX: Support relative paths", "testing/files", []string{"testing", "files"}},
// 		{"UNIX: Support trailing forward slash", "testing/files/", []string{"testing", "files"}},

// 		// wind tests
// 		{"WIND: Support absolute paths", "C:\\testing\\files", []string{"testing", "files"}},
// 		{"WIND: Support relative paths", "\\testing\\files", []string{"testing", "files"}},
// 		{"WIND: Support trailing back slash", "\\testing\\files\\", []string{"testing", "files"}},
// 	}

// 	for _, test := range testCases {
// 		name := helpers.Format("%s", test.message)

// 		t.Run(name, func(t *testing.T) {
// 			result := splitPath(test.path)

// 			helpers.AssertEqualsObject(t, test.expected, result)
// 		})
// 	}
// }
//
//

func TestSplitDirectories(t *testing.T) {
	testCases := []struct {
		desc     string
		filePath string
		dirs     []string
		filename string
	}{
		{
			desc:     "Multiple dirs, name without `.`",
			filePath: "asd/asd/a",
			dirs:     []string{"asd", "asd"},
			filename: "a",
		},
		{
			desc:     "Multiple dirs, name with `.`",
			filePath: "asd/asda/a.text",
			dirs:     []string{"asd", "asda"},
			filename: "a.text",
		},
		{
			desc:     "Pure Directory",
			filePath: "asd/",
			dirs:     []string{"asd"},
			filename: "",
		},
		{
			desc:     "Pure Filename",
			filePath: "asd",
			dirs:     []string{},
			filename: "asd",
		},
		// {
		// 	desc:     "",
		// 	filePath: "",
		// 	dirs:     []string{},
		// 	filename: "",
		// },
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			dirs, filename := SplitDirectories(tC.filePath)

			helpers.AssertEqualsObject(t, tC.dirs, dirs)
			helpers.AssertEquals(t, tC.filename, filename)
		})
	}
}

func TestSplitFilePath(t *testing.T) {
	// when splitting the filename, the cases we care about are:

	testCases := []struct {
		message string

		path     string
		filename string
	}{
		{"./test.txt1", "./", "test.txt1"},
		{"./test.txt.txt2", "./", "test.txt.txt2"},
		{"../test.txt3", "../", "test.txt3"},
		{"/test.txt4", "/", "test.txt4"},
		{"/test", "/", "test"},
		{"asd.md6", "", "asd.md6"},
		{"../custom_skins/custom_skins", "../custom_skins/", "custom_skins"},
		{"../dd/custom_skins/cas.cs", "../dd/custom_skins/", "cas.cs"},
	}

	for _, test := range testCases {

		filename := test.filename

		t.Run(" "+test.message, func(t *testing.T) {
			path, fn := SplitFilePath(test.message)

			helpers.AssertEquals(t, test.path, path)
			helpers.AssertEquals(t, filename, fn)
		})
	}
}

func TestSplitFileName(t *testing.T) {
	testCases := []struct {
		message string

		path     string
		filename string
		suffix   string
	}{
		{"Test splitting a file name with nothing else to worry about", "asd.md", "asd", "md"},
		{"Test splitting a file name with no specified type in it", "/test", "/test", ""},
		{"Test splitting a file name with no actual name it in", ".txt", "", "txt"},

		{"Test splitting a file name with no actual name or type", ".", "", ""},
	}

	for _, test := range testCases {
		t.Run(test.message, func(t *testing.T) {
			fn, sf := SplitFileName(test.path)

			helpers.AssertEquals(t, test.filename, fn)
			helpers.AssertEquals(t, test.suffix, sf)
		})
	}
}

func TestJoinDirs(t *testing.T) {
	testCases := []struct {
		desc string
		dirs []string
		path string
	}{
		{
			desc: "Joinging adds a directory sign at the end",
			dirs: []string{"asd", "asd"},
			path: "asd/asd/",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			path := JoinDirs(tC.dirs...)

			helpers.AssertEquals(t, tC.path, path)
		})
	}
}
