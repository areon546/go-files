package table

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
		records []Row

		hasHeaders    bool
		headersEdited bool
	}
)

// assumptions:
//

// Creates a new table with the specified number of columns and rows.
// NOTE: rows refers to the number of entries in the table, and excludes the Header row.
func NewTable(cols, rows int) *Table {
	t := table{records: makeRows(rows, cols), headers: *NewRow(cols), headersEdited: true}
	return &Table{t}
}

// Returns whether the Table has had all columns populated with a header.
func (t *table) HasHeaders() bool {
	if t.headersEdited {
		// recalculate hasHeaders
		missing := t.MissingHeaders()
		t.hasHeaders = len(missing) == 0
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
	return t.headers.size
}

// Getters

// Returns the array of Rows in the table
func (t *table) Records() []Row {
	return t.records
}

// Returns the Header Row
func (t *table) Headers() (Row, error) {
	if t.HasHeaders() {
		return t.headers, nil
	}
	return *EmptyRow(), ErrHeaderMissing
}

// Specific Getters

// Returns the specified Row across the records
func (t *table) Record(i int) (Row, error) {
	if i < 0 || i > len(t.records) {
		return *NewRow(0), ErrOutOfBounds
	}
	return t.records[i], nil
}

// Returns the specified column in the table.
func (t *table) Col(i int) (string, Row, error) {
	return "", *EmptyRow(), nil
}

// Returns the Header of the specified column
func (t *table) Header(i int) (Cell, error) {
	return *NewCell(""), nil
}

// Setters

// Appends a Row to the Records
func (t *table) AddRecord(r Row) {
	t.records = append(t.records, r)
}

// Adds a column to the table.
// Set the header parameter to "" to leave it empty.
func (t *table) AddCol(header string) {
	for _, row := range t.records {
		row.Lengthen(1)
	}

	// Assign header
	t.headers.Lengthen(1)
	t.headers.Set(t.headers.size-1, header)
}

// Sets the value of a Header based on the specified column index.
// NOTE: If the error value returned is ErrCellPopulated, it has overwritten the contents.
func (t *table) SetHeader(index int, newHeader string) error {
	t.headersEdited = true
	return t.headers.Set(index, newHeader)
}
