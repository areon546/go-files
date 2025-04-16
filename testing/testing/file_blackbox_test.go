package files_test

import (
	"testing"

	"github.com/areon546/go-files/files"
)

func TestOpenFile(t *testing.T) {

	t.Run("New File", func(t *testing.T) {
		files.OpenFile("files/test.txt")

	})

	t.Run("Non-existant File", func(t *testing.T) {

	})

	t.Run("New File", func(t *testing.T) {

	})
}
