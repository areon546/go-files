package table

// Contract:
// I will contain a list of Cells with a Set length.
// If you want a specific cell, you give me the number and I give what is inside of it.
// If you want to edit a specific cell, give me the value and I will change it for you.
// - If you want to edit a cell that doesn't exist, I will tell you that it doesn't exist. ErrIndexOutOfBounds
// You cannot make me longer or shorter.
type (
	Row struct {
		row
	}
	row struct {
		cells  []Cell
		maxLen int
	}
)

func NewRow(length int) *Row {
	return &Row{row{cells: make([]Cell, length), maxLen: length}}
}

// Why does this not return an out of bounds error?
func (r *row) Set(index int, value string) error {
	if index < r.maxLen {
		r.cells[index] = *NewCell(value)
		return nil
	}
	return ErrEndOfRow
}

func (r *row) Get(index int) string {
	return r.cells[index].String()
}

func (r *row) Lengthen(increase int) {
	r.maxLen += increase
}

func (r *row) Size() int {
	return len(r.cells)
}

func (r *row) String() string {
	return format("maxLen: %d, cells: %s", r.maxLen, r.cells)
}

// Creates [length] rows
func makeRows(numberRows, rowLength int) []Row {
	rows := make([]Row, 0)
	for range numberRows {
		rows = append(rows, *NewRow(rowLength))
	}
	return rows
}
