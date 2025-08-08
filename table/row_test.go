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
		helpers.AssertNoError(t, err)

		val, err := r.Get(2)
		helpers.AssertEquals(t, val, firstVal)
		helpers.AssertNoError(t, err)
	})

	t.Run("Overwrite Value", func(t *testing.T) {
		err := r.Set(2, secondVal)
		helpers.AssertError(t, err, ErrCellPopulated)

		val, err := r.Get(2)
		helpers.AssertEquals(t, val, secondVal)
		helpers.AssertNoError(t, err)
	})
}

func TestRowGet(t *testing.T) {
	r := NewRow(3)
	// Test Get for a value that has been Set
	// Test Get for a value that has not been set
	//

	t.Run("Row value not yet set", func(t *testing.T) {
		val, err := r.Get(0)
		helpers.AssertEquals(t, val, "")
		helpers.AssertNoError(t, err)
	})

	t.Run("Row value now set", func(t *testing.T) {
		err := r.Set(0, "ABC")
		helpers.AssertNoError(t, err)

		val, err := r.Get(0)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, val, "ABC")
	})
}

func TestRowLengthen(t *testing.T) {
	t.Run("If you make it longer, new cells are set to \"\"", func(t *testing.T) {
		r := NewRow(3)
		r.Lengthen(1)
		helpers.AssertEqualsInt(t, 4, r.Size())

		s, err := r.Get(3)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, "", s)
	})

	t.Run("If you make it shorter, you cannot access the originally last couple values", func(t *testing.T) {
		r := NewRow(3)
		err := r.Set(2, "asd")
		helpers.AssertNoError(t, err)

		r.Lengthen(-1)
		helpers.AssertEqualsInt(t, 2, r.Size())

		s, err := r.Get(2)
		helpers.AssertError(t, err, ErrOutOfBounds)
		helpers.AssertEquals(t, "", s)

		r.Lengthen(1)
		s, err = r.Get(2)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, "", s)
	})

	t.Run("At length 0, you can make it longer", func(t *testing.T) {
		r := NewRow(0)
		helpers.AssertEqualsInt(t, 0, r.Size())

		r.Lengthen(1)
		helpers.AssertEqualsInt(t, 1, r.Size())

		val, err := r.Get(0)
		helpers.AssertEquals(t, "", val)
		helpers.AssertNoError(t, err)
	})
}

// Test that you get the row back
// test thta when you shorten it, you get the appropriate elements regardless
func TestRowCells(t *testing.T) {
	r := NewRow(5)

	r.Set(0, "0")
	r.Set(1, "1")
	r.Set(2, "2")
	r.Set(3, "3")
	r.Set(4, "4")

	t.Run("Get Access to Row object", func(t *testing.T) {
		row := r.Cells()
		eRow := []Cell{*NewCell("0"), *NewCell("1"), *NewCell("2"), *NewCell("3"), *NewCell("4")}

		helpers.AssertEqualsObject(t, eRow, row)
	})

	t.Run("Shorter = cannot access last few values. ", func(t *testing.T) {
		r.Lengthen(-2)

		row := r.Cells()
		eRow := []Cell{*NewCell("0"), *NewCell("1"), *NewCell("2")}

		helpers.AssertEqualsObject(t, eRow, row)
	})
}

func TestRowSize(t *testing.T) {
	r := NewRow(3)
	// Test Size of row can only be adjusted by (1) method
	t.Run("Size correlates to preset value", func(t *testing.T) {
		helpers.AssertEqualsInt(t, 3, r.Size())
	})
	// Test access to values based on make lenght
	t.Run("If you change size, it is reflected", func(t *testing.T) {
		r.Lengthen(-1)
		helpers.AssertEqualsInt(t, 2, r.Size())
	})
}

func TestRowString(t *testing.T) {
	// TODO: unsure what to have a an ideal string representation of a row
}

func TestJoin(t *testing.T) {
	// TODO: test if it joins properly into a CSV line
}
