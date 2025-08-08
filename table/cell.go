package table

import "reflect"

// Contract:
// I contain a string.
// If you give me a string, I will then give it back to you.
// If you want to change my string, I will allow you.
type (
	Cell struct {
		cell
	}
	cell struct {
		value string
	}
)

func NewCell(value string) *Cell {
	return &Cell{cell{value: value}}
}

func EmptyCell() *Cell {
	return NewCell("")
}

func (c *cell) String() string {
	return c.value
}

func (c *cell) Set(new string) error {
	cellHadValue := !c.IsEmpty()
	c.value = new

	if cellHadValue {
		return ErrCellPopulated
	}
	return nil
}

// Returns whether the Cell contains "" or not.
func (c *cell) IsEmpty() bool {
	return reflect.DeepEqual(c.value, "")
}
