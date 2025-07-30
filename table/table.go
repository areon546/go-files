package table

import "github.com/areon546/go-helpers/helpers"

type TableConverter func(t table) string

// Contract:
// I can make a table for you, with or without headers.
// I can add rows and columns to my table.
// I can give you the value at any specific cell, if you give my it's location.
// I can assign headers to specified columns.
// - I will tell you if there are still columns without headers.
// If you want to adjust my size, you have to give me values to populate the remaining elements.
type (
	Table struct{ table }
	table struct {
		headers Row
		records []Row

		hasHeaders    bool
		headersEdited bool
	}
)

// assumptions:
//

// Creates a new table with the specified number of columns and rows.
// NOTE: rows refers to the number of entries in the table, and excludes the Header row.
func NewTable(cols int) *Table {
	t := table{records: makeRows(0, cols), headers: *NewRow(cols), headersEdited: true}
	return &Table{t}
}

// Returns whether the Table has had all columns populated with a header.
func (t *table) HasHeaders() bool {
	if t.headersEdited {
		// recalculate hasHeaders
		missing := t.MissingHeaders()
		noMissingElements := len(missing) == 0
		atLeastOneCol := t.Cols() != 0

		t.hasHeaders = noMissingElements && atLeastOneCol
		t.headersEdited = false
	}

	return t.hasHeaders
}

// Optional function.
// Returns the indexes of columns without headers.
func (t *table) MissingHeaders() []int {
	headers := []int{}

	for i, val := range t.headers.cells {
		// each loop, check if header is present
		if val.IsEmpty() {
			headers = append(headers, i)
		}
	}

	return headers
}

// Returns the number of records the table has been populated with.
func (t *table) Entries() int {
	return len(t.records)
}

// Returns the number of columns the table has.
func (t *table) Cols() int {
	return t.headers.Size()
}

// Getters

// Returns the array of Rows in the table
func (t *table) Records() []Row {
	return t.records
}

// Returns the Header Row.
// NOTE: If there isn't a header in the table, it will return an empty row of length zero, and ErrHeaderMissing
func (t *table) Headers() (Row, error) {
	if t.HasHeaders() {
		return t.headers, nil
	}
	return *EmptyRow(), ErrHeaderMissing
}

// Specific Getters

// Returns the specified Row across the records
func (t *table) Record(i int) (Row, error) {
	if indexWithinBounds(i, t.Entries()) {
		return t.records[i], nil
	}
	return *NewRow(0), ErrOutOfBounds
}

// Returns the specified column in the table.
func (t *table) Col(i int) (string, Row, error) {
	return "", *EmptyRow(), nil
}

// Returns the Header of the specified column
func (t *table) Header(i int) (Cell, error) {
	if t.HasHeaders() {
		return t.headers.cells[i], nil
	}

	return *NewCell(""), nil
}

// Setters

// Appends a Row to the Records
func (t *table) AddRecord(r Row) {
	t.records = append(t.records, r)
}

// Adds a column to the table.
// Set the header parameter to "" to leave it empty.
func (t *table) AddCol(header string, colValues Row) error {
	// check if values.Size == table.Size,
	numRows := t.Cols()
	inputColumnSize := colValues.Size()

	tableHasASize := t.Cols() > 0
	incompatibleColumnSize := numRows != inputColumnSize
	if tableHasASize && incompatibleColumnSize {
		return ErrIncompatibleSize
	}

	// No error, thus want to actually add columns.
	for recordIndex, row := range t.records {
		row.Lengthen(1)

		colEntry, err := colValues.Get(recordIndex)
		helpers.Handle(err)
		err = row.Set(row.Size()-1, colEntry)
		helpers.Handle(err)
	}

	// Assign header
	lastIndex := t.headers.Size()
	t.headers.Lengthen(1)
	t.headers.Set(lastIndex, header)

	return nil
}

// Sets the value of a Header based on the specified column index.
// NOTE: If the error value returned is ErrCellPopulated, it has overwritten the contents.
func (t *table) SetHeader(index int, newHeader string) error {
	t.headersEdited = true
	return t.headers.Set(index, newHeader)
}
