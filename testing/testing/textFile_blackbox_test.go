package files_test

import (
	"testing"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

func TestNewTextFile(t *testing.T) {

	t.Run("New Test file", func(t *testing.T) {
		txt := files.NewTextFile("/testing-files/test.txt", "")

		want := helpers.StringToBytes("[[dragons are the best]]")
		get := txt.ReadFile()

		helpers.AssertEqualsObject(t, want, get)
	})

}
