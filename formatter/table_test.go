package formatter

import (
	"testing"
)

var (
	rows      int   = 1
	cols      int   = 5
	headers   row   = setupHeaderRow()
	rowValues []row = []row{}
	t         Table = Table{table{headers: headers, rows: rowValues}}
)

func setupHeaderRow() row {
	r := *NewRow(cols)

	return r
}

func TestNewTable(t *testing.T) {
	cols, rows := 1, 5
	want := &Table{table{headers: *NewRow(cols), rows: makeRows(rows, cols)}}
	get := NewTable(cols, rows)

	assertEqualsObject(t, want, get)
}

func TestRows(t *testing.T) {

}
