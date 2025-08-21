package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

var existingFile *TextFile

func init() {
	helpers.Print("Initialising textfile_test")
	defer helpers.Print("Finished initialising textfile_test\n")
}

func TestNewTextFile(t *testing.T) {
	t.Run("Open Existing file", func(t *testing.T) {
		txt := NewTextFile("files/textfile-test-1.txt")

		want := []string{"[dragons are cool]", ""}
		get, err := txt.ReadContents()
		helpers.AssertNoError(t, err)

		// test content
		helpers.AssertEqualsInt(t, 2, len(get))
		helpers.AssertEqualsObject(t, want, get)
	})

	t.Run("Open file that doesn't exist", func(t *testing.T) {
		// txt := NewTextFile("files/test2/txt")
		// want := []string{}
		// get, err := txt.ReadContents()

		// helpers.AssertError(t, err, ErrNoFileOrDirectory) // NOTE: See file.ReadContents for relevant error to use
		// helpers.AssertEqualsObject(t, want, get)
		//
		// print(txt)

		// helpers.AssertEqualsInt(t, 0, 1)
	})
}

func TestConentsTextFile(t *testing.T) {
	// TODO: test contents of a text file are: correct
}
