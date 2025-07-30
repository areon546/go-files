package table

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

var (
	recordsToAdd []Row = []Row{
		*setTestRow("0", "A", "Black", "5", "5"),
		*setTestRow("1", "AB", "Black", "5", "5"),
		*setTestRow("2", "ABC", "Black", "5", "5"),
		*setTestRow("3", "ABCD", "Black", "5", "5"),
		*setTestRow("4444", "ABCDE", "Black", "5", "5"),
	}

	globalTable Table = testTable()
	emptyTable  Table = testEmptyTable()
)

func testTable() Table {
	tab := NewTable(5)
	tab.SetHeader(0, "Horse ID")
	tab.SetHeader(1, "Name")
	tab.SetHeader(2, "Colour")
	tab.SetHeader(3, "Pattern")
	tab.SetHeader(4, "Speed")

	for _, record := range recordsToAdd {
		tab.AddRecord(record)
	}
	return *tab
}

func testEmptyTable() Table {
	tab := NewTable(5)
	return *tab
}

func TestNewTable(t *testing.T) {
	cols, rows := 5, 0
	want := &Table{table{headers: *NewRow(cols), records: makeRows(rows, cols), headersEdited: true}}
	tab := NewTable(cols)
	helpers.AssertEqualsObject(t, want, tab)
}

// test HasHeaders
func TestHasHeaders(t *testing.T) {
	// test:
	//   if it recompiles headers
	t.Run("Does it recompile on change", func(t *testing.T) {
		// TODO: TODO
	})

	t.Run("Complete Headers", func(t *testing.T) {
		tab := NewTable(1)

		err := tab.SetHeader(0, "New Header")
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsBool(t, true, tab.HasHeaders())
	})
	t.Run("Incomplete Headers", func(t *testing.T) {
		tab := NewTable(2)

		err := tab.SetHeader(0, "New Header")
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsBool(t, false, tab.HasHeaders())
		tab.MissingHeaders()
	})
	t.Run("Missing Headers", func(t *testing.T) {
		tab := NewTable(2)

		helpers.AssertEqualsBool(t, false, tab.HasHeaders())
	})

	t.Run("Length 0", func(t *testing.T) {
		tab := NewTable(0)

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
		tab := NewTable(1)

		err := tab.SetHeader(0, "New Header")
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsObject(t, []int{}, tab.MissingHeaders())
	})
	t.Run("Incomplete Headers", func(t *testing.T) {
		tab := NewTable(2)

		err := tab.SetHeader(0, "New Header")
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsObject(t, []int{1}, tab.MissingHeaders())
		tab.MissingHeaders()
	})
	t.Run("Missing Headers", func(t *testing.T) {
		tab := NewTable(2)

		helpers.AssertEqualsObject(t, []int{0, 1}, tab.MissingHeaders())
	})

	t.Run("Length 0", func(t *testing.T) {
		tab := NewTable(0)

		helpers.AssertEqualsObject(t, []int{}, tab.MissingHeaders())
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
	tab := globalTable
	rowLen := tab.Cols()

	helpers.AssertEqualsInt(t, 5, rowLen)
	helpers.AssertEqualsObject(t, recordsToAdd, tab.Records())
}

// TestHeaders
func TestHeaders(t *testing.T) {
	t.Run("No Headers", func(t *testing.T) {
		tab := emptyTable
		headers, err := tab.Headers()
		helpers.AssertError(t, err, ErrHeaderMissing)
		helpers.AssertEqualsObject(t, *EmptyRow(), headers)
	})

	t.Run("Has Headers", func(t *testing.T) {
		tab := globalTable
		headers, err := tab.Headers()
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsObject(t, tab.headers, headers)
	})
}

//

// TestRecord
func TestRecord(t *testing.T) {
}

// TestCol
func TestCol(t *testing.T) {
}

// TestHeader
func TestHeader(t *testing.T) {
}

//

// AddRow
func TestAddRecord(t *testing.T) {
	// what does adding a record do?
	// it increases the size
	// it creates a New thing at the end (check for err out of bounds b4 and After)
}

// AddCol
func TestAddCol(t *testing.T) {
	tab := NewTable(0)
	helpers.AssertEqualsInt(t, 0, tab.Cols())

	// NOTE: commented out because currently crashing tests
	row := NewRow(5)
	err := tab.AddCol("ID", *row)
	helpers.AssertNoError(t, err)

	tooLongRow := NewRow(7)
	err = tab.AddCol("Speed", *tooLongRow)
	helpers.AssertError(t, err, ErrIncompatibleSize)
}

// SetHeader
func TestSetHeader(t *testing.T) {
	t.Run("No Headers in table", func(t *testing.T) {
		tab := NewTable(0)

		hd, err := tab.Header(0)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, "", hd.String())
	})

	t.Run("Headers in table", func(t *testing.T) {
		tab := globalTable

		hd, err := tab.Header(0)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, "Horse ID", hd.String())
	})
}

func setTestRow(v1, v2, v3, v4, v5 string) *Row {
	r := NewRow(5)
	r.Set(0, v1)
	r.Set(1, v2)
	r.Set(2, v3)
	r.Set(3, v4)
	r.Set(4, v5)

	return r
}
