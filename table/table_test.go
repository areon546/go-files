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
		tab.AddRecord(&record)
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

// Width
func TestWidth(t *testing.T) {
	tab := globalTable
	helpers.AssertEqualsInt(t, 5, tab.Width())
}

func TestEntries(t *testing.T) {
	tab := globalTable
	helpers.AssertEqualsInt(t, 5, tab.Entries())
}

func TestIsCompatible(t *testing.T) {
	tab := emptyTable
	record := NewRow(5)

	helpers.AssertEqualsBool(t, true, tab.IsCompatible(record))
}

func TestWiden(t *testing.T) {
}

//

// TestRecords
func TestRecords(t *testing.T) {
	tab := globalTable
	rowLen := tab.Width()

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
	t.Run("accessing within bounds", func(t *testing.T) {
		// TODO:
	})

	t.Run("accessing outside of bounds", func(t *testing.T) {
		// TODO:
	})
}

// IndexOf ()
func TestIndexOf(t *testing.T) {
	tab := globalTable

	t.Run("Valid header", func(t *testing.T) {
		header := "Horse ID"
		headerI := 0
		index := tab.IndexOf(header)

		helpers.AssertEqualsInt(t, headerI, index)

		headerVal, err := tab.Header(index)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, header, headerVal)
	})
	t.Run("Invalid header", func(t *testing.T) {
		index := tab.IndexOf("")
		helpers.AssertEqualsInt(t, -1, index)
	})
}

// Cell
func TestCell(t *testing.T) {
	tab := globalTable

	t.Run("Valid index", func(t *testing.T) {
		val, err := tab.Cell(0, 0)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, "0", val)
	})

	t.Run("Invalid index", func(t *testing.T) {
		val, err := tab.Cell(-1, -1)
		helpers.AssertError(t, err, ErrOutOfBounds)
		helpers.AssertEquals(t, "", val)
	})
}

// TestCol
func TestCol(t *testing.T) {
	t.Run("accessing within bounds", func(t *testing.T) {
		// TODO:
	})

	t.Run("accessing outside of bounds", func(t *testing.T) {
		// TODO:
	})
}

// TestHeader
func TestHeader(t *testing.T) {
	t.Run("accessing within bounds", func(t *testing.T) {
		// TODO:
	})

	t.Run("accessing outside of bounds", func(t *testing.T) {
		// TODO:
	})
}

//

// AddRecord
func TestAddRecord(t *testing.T) {
	record := NewRow(2)
	record.Set(0, "length")
	record.Set(1, "width")

	tab := NewTable(2)
	helpers.AssertEqualsInt(t, 0, tab.Entries())

	// ErrOutOfBounds when trying to access before populated
	row0, err := tab.Record(0)
	helpers.AssertError(t, err, ErrOutOfBounds)
	_, err = row0.Get(0)
	helpers.AssertError(t, err, ErrOutOfBounds)

	tab.AddRecord(record)
	helpers.AssertEqualsInt(t, 1, tab.Entries())

	// No ErrOutOfBounds after population
	row0, err = tab.Record(0)
	helpers.AssertNoError(t, err)
	row0Val0, err := row0.Get(0)
	helpers.AssertNoError(t, err)
	helpers.AssertEquals(t, "length", row0Val0)

	// what does adding a record do?
	// it increases the size
	// it creates a New thing at the end (check for err out of bounds b4 and After)
}

// AddRecords
func TestAddRecords(t *testing.T) {
}

// AddCol
func TestAddCol(t *testing.T) {
	tab := NewTable(0)
	helpers.AssertEqualsInt(t, 0, tab.Width())

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
		helpers.AssertEquals(t, "", hd)
	})

	t.Run("Headers in table", func(t *testing.T) {
		tab := globalTable

		hd, err := tab.Header(0)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, "Horse ID", hd)
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
