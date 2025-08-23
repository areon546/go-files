package table

import "iter"

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

// Returns ErrOutOfBounds if index is too big or small.
func (r *row) Set(index int, value string) error {
	if indexWithinBounds(index, r.size) {
		return r.cells[index].Set(value)
	}
	return ErrOutOfBounds
}

// Returns ErrOutOfBounds if given index is larger than the size of the Row.
func (r *row) Get(index int) (string, error) {
	if indexWithinBounds(index, r.size) {
		return r.cells[index].String(), nil
	}
	return "", ErrOutOfBounds // Standard result if index is out of bounds
}

func (r *row) Lengthen(increaseBy int) {
	if increaseBy == 0 {
		return
	}

	r.size += increaseBy

	needToAddElements := r.size > len(r.cells)
	if needToAddElements {
		elementsToAdd := r.size - len(r.cells)

		if elementsToAdd > 0 {
			// add to length of r.cells
			r.cells = append(r.cells, make([]Cell, elementsToAdd)...)
		} else if elementsToAdd <= 0 {
			// NOTE: elementsToAdd should never be zero nor lower
			panic("elementsToAdd is 0 or lower")
		}
	}

	// Need to reset values at tail of row if shortening.
	// It doesn't make sense if when we shorten the row, it will keep values at the tail of the slice.
	if increaseBy < 0 {
		for indexToReset := len(r.cells) + increaseBy; indexToReset < len(r.cells); indexToReset++ {
			r.cells[indexToReset].Set("")
		}
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

func (r *row) Join(beforeRow, cellPre, inbetweenCells, cellSuff, afterRow string) string {
	str := ""

	for i, cell := range r.cells {
		str += cellPre + cell.value + cellSuff

		// DO NOT add `inbetweenCells` for the very last cell
		if i < r.size-1 { // -1 because `i` is an index, while `size` is a length
			str += inbetweenCells
		}
	}
	return beforeRow + str + afterRow
}

// Iteration
func (r *row) Iter() iter.Seq2[int, Cell] {
	return func(yield func(int, Cell) bool) {
		for i, cell := range r.cells {
			if !yield(i, cell) {
				return
			}
		}
	}
}

// Creates [length] rows
func makeRows(numberRows, rowLength int) []Row {
	rows := make([]Row, 0)
	for range numberRows {
		rows = append(rows, *NewRow(rowLength))
	}
	return rows
}
