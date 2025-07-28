package table

// Contract:
// I will contain a list of Cells with a Set length.
// If you want a specific cell, you give me the number and I give what is inside of it.
// If you want to edit a specific cell, give me the value and I will change it for you.
// - If you want to edit a cell that doesn't exist, I will tell you that it doesn't exist. ErrIndexOutOfBounds
// You can make me longer or shorter.
// Be familiar, use index notation where possible.
type (
	Row struct {
		row
	}
	row struct {
		cells []Cell
		size  int
	}
)

func NewRow(length int) *Row {
	return &Row{row{cells: make([]Cell, length), size: length}}
}

// Default Row, with no extra information.
func EmptyRow() *Row {
	return NewRow(0)
}

// Why does this not return an out of bounds error?
func (r *row) Set(index int, value string) error {
	if index < r.size {
		r.cells[index] = *NewCell(value)
		return nil
	}
	return ErrOutOfBounds
}

func (r *row) Get(index int) (string, error) {
	if index < r.size {
		return r.cells[index].String(), nil
	}
	return "", ErrOutOfBounds // Standard result if index is out of bounds
}

func (r *row) Lengthen(increaseBy int) {
	if increaseBy == 0 {
		return
	}

	r.size += increaseBy

	// NOTE: Do not have a increaseBy<0 branch because if the length gets shorter
	// we don't actually need to adjust the length of the slice, just the datastructure's error lenght.
	lengthToIncrease := len(r.cells) - increaseBy
	if increaseBy > 0 && lengthToIncrease > 0 {
		r.cells = append(r.cells, make([]Cell, lengthToIncrease)...)
	}
}

func (r *row) Cells() []Cell {
	return r.cells[0:r.size]
}

func (r *row) Size() int {
	return r.size
}

func (r *row) String() string {
	return format("size: %d, cells: %s", r.size, r.cells)
}

// Creates [length] rows
func makeRows(numberRows, rowLength int) []Row {
	rows := make([]Row, 0)
	for range numberRows {
		rows = append(rows, *NewRow(rowLength))
	}
	return rows
}
