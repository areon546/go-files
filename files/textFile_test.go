package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func TestNewTextFile(t *testing.T) {
	t.Run("Open Existing file", func(t *testing.T) {
		txt := NewTextFile("files/textfile-test-1.txt")

		want := "[dragons are cool]"
		get := txt.ReadContents()[0]
		print(get)

		// test content
		helpers.AssertEqualsObject(t, want, get)
	})

	t.Run("Open file that doesn't exist", func(t *testing.T) {
		txt := NewTextFile("files/test2/txt")

		print(txt)

		// helpers.AssertEqualsInt(t, 0, 1)
	})
}

func TestConentsTextFile(t *testing.T) {
	// TODO: test contents of a text file are: correct
}
