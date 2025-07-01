package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func TestNewTextFile(t *testing.T) {
	t.Run("Open Existing file", func(t *testing.T) {
		txt := NewTextFile("files/test.txt")

		want := "[[dragons are the best]]"
		get := txt.ReadFile()[0]

		// test content
		helpers.AssertEqualsObject(t, want, get)
	})

	t.Run("Open file that doesn't exist", func(t *testing.T) {
		// txt := files.NewTextFile("files/test2/txt")
	})
}

func TestConentsTextFile(t *testing.T) {
}
