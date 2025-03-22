package files_test

import (
	. "github/com/areon546/go-files/files"
	"testing"
)

func TestNewTextFile(t *testing.T) {

	t.Run("New Test file", func(t *testing.T) {
		txt := NewTextFile()

		print(txt.Read())
	})

}
