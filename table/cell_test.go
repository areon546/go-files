package table

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

// Contract:
// I contain a string.
// If you give me a string, I will then give it back to you.
// If you want to change my string, I will allow you.

// Test the cell.String function returns the correct output
func TestCellString(t *testing.T) {
	val := "hello_world"
	c := NewCell("hello_world")

	helpers.AssertEquals(t, val, c.String())
}

func TestCellSet(t *testing.T) {
	newVal := "NEW CELL VALUE"
	c := NewCell("Initial Value")
	c.Set(newVal)

	helpers.AssertEquals(t, newVal, c.String())
}
