package table

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

var tab *Table = testTable()

func testTable() *Table {
	rows := 5
	cols := 5
	tab := &Table{table: table{headers: *NewRow(cols), rows: makeRows(rows, cols)}}

	return tab
}

func TestNewTable(t *testing.T) {
	cols, rows := 1, 5
	want := &Table{table{headers: *NewRow(cols), rows: makeRows(rows, cols)}}
	get := NewTable(cols, rows, false)

	helpers.AssertEqualsObject(t, want, get)
}

// func TestRows(t *testing.T) {
// 	helpers.AssertEquals(t, "", "a")
// }
//
// func TestCols(t *testing.T) {
// 	helpers.AssertEquals(t, "", "a")
// }
//
// func TestAddRow(t *testing.T) {
// 	helpers.AssertEquals(t, "", "a")
// }
//
// func TestAddCol(t *testing.T) {
// 	helpers.AssertEquals(t, "", "a")
// }
//
// func TestGetRow(t *testing.T) {
// 	helpers.AssertEquals(t, "", "a")
// }
//
// func TestGetCol(t *testing.T) {
// 	helpers.AssertEquals(t, "", "a")
// }
//
// func TestSetHeader(t *testing.T) {
// 	helpers.AssertEquals(t, "", "a")
// }
