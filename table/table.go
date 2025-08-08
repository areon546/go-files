package table

import (
	"iter"
	"reflect"

	"github.com/areon546/go-helpers/helpers"
)

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

func NewPopulatedTable(cols int, records *[]Row) *Table {
	t := NewTable(cols)
	t.AddRecords(records)
	return t
}

func EmptyTable() *Table {
	return NewTable(0)
}

func (t *table) String() string {
	headers := t.headers.String()

	rows := ""
	for _, row := range t.Iter() {
		rows += row.String() + "\n"
	}

	return helpers.Format("Headers: %s, \nRows: \n%s", headers, rows)
}

// Lets you for x,y := range table.Iter() and loop across the rows.
func (t *table) Iter() iter.Seq2[int, Row] {
	return func(yield func(int, Row) bool) {
		for i, row := range t.records {
			if !yield(i, row) {
				return
			}
		}
	}
}

// Returns whether the Table has had all columns populated with a header.
func (t *table) HasHeaders() bool {
	if t.headersEdited {
		// recalculate hasHeaders
		missing := t.MissingHeaders()
		noMissingElements := len(missing) == 0
		atLeastOneCol := t.Width() != 0

		t.hasHeaders = noMissingElements && atLeastOneCol
		t.headersEdited = false
	}

	return t.hasHeaders
}

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
func (t *table) Width() int {
	return t.headers.Size()
}

// Returns whether the given Record is compatible with the instance of the table due to length, or not.
func (t *table) IsCompatible(record *Row) bool {
	return record.Size() == t.Width()
}

// Getters

// Returns the array of Rows in the table
func (t *table) Records() []Row {
	return t.records
}

// Returns the Header Row.
// NOTE: If there isn't a header in the table, it will return an empty row of length zero, and ErrHeaderMissing
func (t *table) Headers() (Row, error) { // Pass by value because I don't want a user to be able to edit the values of the headers separately.
	if t.HasHeaders() {
		return t.headers, nil
	}
	return *EmptyRow(), ErrHeaderMissing
}

// Specific Getters

// Returns the specified Row across the records
func (t *table) Record(i int) (Row, error) { // Pass by value because I don't want the use editing values of the reocrds.
	if indexWithinBounds(i, t.Entries()) {
		return t.records[i], nil
	}
	return *NewRow(0), ErrOutOfBounds
}

func (t *table) Cell(recordIndex, column int) (string, error) {
	record, err := t.Record(recordIndex)
	if err != nil {
		return EmptyCell().value, err
	}

	return record.Get(column)
}

// Returns the specified column in the table.
func (t *table) Col(i int) (header string, colVals *Row, err error) {
	return "", EmptyRow(), nil // TODO:
}

// Returns the Header of the specified column
func (t *table) Header(i int) (string, error) {
	if t.HasHeaders() {
		return t.headers.cells[i].String(), nil
	}

	return EmptyCell().value, nil
}

// Returns the index of a given heading in the table.
// Returns the index of the first encountered heading if there are multiple to choose from.
func (t *table) IndexOf(heading string) int {
	for i, cell := range t.headers.Iter() {
		if reflect.DeepEqual(heading, cell.String()) {
			return i
		}
	}
	return -1
}

// Adders

// Appends a Row to the Records
func (t *table) AddRecord(r *Row) error {
	if !t.IsCompatible(r) {
		return ErrIncompatibleSize
	}
	t.records = append(t.records, *r)
	return nil
}

// Adds all the specified records, one by one.
// If ANY of the specified records are incompatible, it will cancel the entire operation and return an error.
func (t *table) AddRecords(records *[]Row) error {
	compatibleRows := []Row{}

	for _, record := range *records {
		if !t.IsCompatible(&record) {
			return ErrIncompatibleSize
		}

		compatibleRows = append(compatibleRows, record)

	}

	t.records = append(t.records, compatibleRows...)
	return nil
}

// Increases the width of the entire table by 1.
// Another way to add columns, however without needing to prepopulate them with values.
func (t *table) Widen(l int) {
	t.headers.Lengthen(l)

	for _, row := range t.Iter() {
		row.Lengthen(l)
	}
}

// Adds a column to the table.
// Set the header parameter to "" to leave it empty.
func (t *table) AddCol(header string, colValues Row) error {
	// check if values.Size == table.Size,
	numRows := t.Entries()
	inputColumnSize := colValues.Size()

	tableHasASize := t.Width() > 0
	incompatibleColumnSize := numRows != inputColumnSize
	if tableHasASize && incompatibleColumnSize {
		return ErrIncompatibleSize
	}

	lastIndex := t.headers.Size()
	t.Widen(1)

	// Assign header
	_ = t.headers.Set(lastIndex, header) // Ignoring error because guaranteed to not fail.

	// No error, thus want to actually add columns.
	for recordIndex, row := range t.records {
		colEntry, err := colValues.Get(recordIndex)
		helpers.Handle(err)
		err = row.Set(row.Size()-1, colEntry)
		helpers.Handle(err)
	}

	return nil
}

// Setters

func (t *table) SetHeaders(headers *Row) error {
	if !t.IsCompatible(headers) {
		return ErrIncompatibleSize
	}

	t.headers = *headers
	return nil
}

// Sets the value of a Header based on the specified column index.
// NOTE: If the error value returned is ErrCellPopulated, it has overwritten the contents.
func (t *table) SetHeader(index int, newHeader string) error {
	t.headersEdited = true
	return t.headers.Set(index, newHeader)
}
