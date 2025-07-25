package table

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func TestNewRow(t *testing.T) {
	l := 5
	want := &Row{row{cells: make([]Cell, l), size: l}}
	get := NewRow(l)

	helpers.AssertEqualsObject(t, want, get)
}

func TestRowSet(t *testing.T) {
	r := NewRow(3)
	firstVal := "First Element"
	secondVal := "Actually, it is the third element"

	t.Run("Index out of bounds", func(t *testing.T) {
		err := r.Set(3, "Alients out of bounds")

		helpers.AssertError(t, err, ErrOutOfBounds)
	})

	t.Run("Set Value", func(t *testing.T) {
		err := r.Set(2, firstVal)

		helpers.AssertEquals(t, r.Get(2), firstVal)
		helpers.AssertNoError(t, err)
	})

	t.Run("Overwrite Value", func(t *testing.T) {
		err := r.Set(2, secondVal)

		helpers.AssertEquals(t, r.Get(2), secondVal)
		helpers.AssertNoError(t, err)
	})
}

func TestRowGet(t *testing.T) {
	// Test Get for a value that has been Set
	// Test Get for a value that has not been set
	//

	helpers.AssertEquals(t, "", "a")
}

func TestRowSize(t *testing.T) {
	// Test Size of row
	helpers.AssertEquals(t, "", "a")
}
