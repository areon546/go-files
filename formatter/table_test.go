package formatter

import (
	"testing"
)

func TestNewTable(t *testing.T) {
	cols, rows := 1, 5
	want := &Table{table{headers: *NewRow(cols), rows: makeRows(rows, cols)}}
	get := NewTable(cols, rows)

	assertEqualsObject(t, want, get)
}

func TestNewRow(t *testing.T) {
	l := 5
	want := &row{cells: make([]cell, l), maxLen: l}
	get := NewRow(l)

	assertEqualsObject(t, want, get)
}
