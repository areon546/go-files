package table

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func TestNewRow(t *testing.T) {
	l := 5
	want := &Row{row{cells: make([]Cell, l), maxLen: l}}
	get := NewRow(l)

	helpers.AssertEqualsObject(t, want, get)
}

func TestSet(t *testing.T) {
	r := NewRow(3)

	t.Run("Index out of bounds", func(t *testing.T) {
		err := r.Set(3, "Alients out of bounds")

		helpers.AssertError(t, err, ErrOutOfBounds)
	})
}
