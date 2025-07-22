package table

import "testing"

func TestNewRow(t *testing.T) {
	l := 5
	want := &row{cells: make([]cell, l), maxLen: l}
	get := NewRow(l)

	assertEqualsObject(t, want, get)
}

func TestSet(t *testing.T) {
	r := NewRow(3)

	t.Run("Index out of bounds", func(t *testing.T) {
		err := r.Set(3, "Alients out of bounds")

		handleErrorExcept(t, err, ErrOutOfBounds)
	})
}
