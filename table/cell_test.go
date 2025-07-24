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
func TestString(t *testing.T) {
	val := "hello_world"
	c := NewCell("hello_world")

	helpers.AssertEquals(t, val, c.String())
}
