package table

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

func (c *cell) String() string {
	return c.value
}

func (c *cell) Set(new string) {
	c.value = new
}
