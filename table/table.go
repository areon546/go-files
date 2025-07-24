package table

import (
	"errors"
)

type TableConverter func(t table) string

// Contract:
// I can make a table for you, with or without headers.
// I can add rows and columns to my table.
// I can give you the value at any specific cell, if you give my it's location.
// I can assign headers to specified columns.
// - I will tell you if there are still columns without headers.
type (
	Table struct{ table }
	table struct {
		headers Row
		rows    []Row

		hasHeaders  bool
		headersFull bool
	}
)

func NewTable(cols, rows int, headers bool) *Table {
	t := table{headers: *NewRow(cols), rows: makeRows(rows, cols), hasHeaders: headers}
	return &Table{t}
}

func (t *table) Rows() int {
	return len(t.rows)
}

func (t *table) Cols() int {
	return t.headers.Size()
}

func (t *table) AddRow(r Row) {
	t.rows = append(t.rows, r)
}

func (t *table) GetRow(i int) (Row, error) {
	if i < 0 || i > t.Rows() {
		return *NewRow(0), ErrOutOfBounds
	}
	return t.rows[i], nil
}

func (t *table) AddCol() {
	t.headers.Lengthen(1)
	for _, row := range t.rows {
		row.Lengthen(1)
	}
}

func (t *table) SetHeader(index int, newHeader string) (err error) {
	err = t.headers.Set(index, newHeader)
	return errors.Join(err, errors.New(": end of headers"))
}
