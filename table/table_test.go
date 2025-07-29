package table

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

var globalTable Table

func TestNewTable(t *testing.T) {
	cols, rows := 5, 5
	want := &Table{table{headers: *NewRow(cols), records: makeRows(rows, cols), headersEdited: true}}
	tab := NewTable(cols, rows)
	helpers.AssertEqualsObject(t, want, tab)

	globalTable = *tab
}

// test HasHeaders
func TestHasHeaders(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")

	// test:
	//   if it recompiles headers
	//   test 100% headers
	//   test 0% headers
	//   test partial headers
	t.Run("Does it recompile on change", func(t *testing.T) {
	})

	t.Run("Complete Headers", func(t *testing.T) {
		tab := NewTable(1, 1)

		err := tab.SetHeader(0, "New Header")
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsBool(t, true, tab.HasHeaders())
	})
	t.Run("Incomplete Headers", func(t *testing.T) {
		tab := NewTable(2, 1)

		err := tab.SetHeader(0, "New Header")
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsBool(t, false, tab.HasHeaders())
		tab.MissingHeaders()
	})
	t.Run("Missing Headers", func(t *testing.T) {
		tab := NewTable(2, 1)

		helpers.AssertEqualsBool(t, false, tab.HasHeaders())
	})
}

// test MissingHeaders
func TestMissingHeaders(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")

	// test 100% headers
	// test 0% headers
	// test partial headers

	t.Run("Complete Headers", func(t *testing.T) {
		tab := NewTable(1, 1)

		err := tab.SetHeader(0, "New Header")
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsObject(t, []int{}, tab.MissingHeaders())
	})
	t.Run("Incomplete Headers", func(t *testing.T) {
		tab := NewTable(2, 1)

		err := tab.SetHeader(0, "New Header")
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsObject(t, []int{1}, tab.MissingHeaders())
		tab.MissingHeaders()
	})
	t.Run("Missing Headers", func(t *testing.T) {
		tab := NewTable(2, 1)

		helpers.AssertEqualsObject(t, []int{0, 1}, tab.MissingHeaders())
	})
}

//

// TestCols
func TestCols(t *testing.T) {
	tab := globalTable
	helpers.AssertEqualsInt(t, 5, tab.Cols())
}

func TestEntries(t *testing.T) {
	tab := globalTable
	helpers.AssertEqualsInt(t, 5, tab.Entries())
}

//

// TestRecords
func TestRecords(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")
}

// TestHeaders
func TestHeaders(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")
}

//

// TestRecord
func TestRecord(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")
}

// TestCol
func TestCol(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")
}

// TestHeader
func TestHeader(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")
}

//

// AddRow
func TestAddRow(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")
}

// AddCol
func TestAddCol(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")
}

// SetHeader
func TestSetHeader(t *testing.T) {
	// helpers.AssertEquals(t, "", "a")
	// check if the row has headers
	// check if the row doesn't have headers
}
