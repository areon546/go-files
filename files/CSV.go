package files

import (
	"errors"
	"strings"

	"github.com/areon546/go-files/table"
	"github.com/areon546/go-helpers/helpers"
)

var (
	errMissingHeaders error = errors.New("headings are missing")
	ErrMissingHeaders error = helpers.WrapError("%w %w", errFiles, errMissingHeaders)
)

// ~~~~~~~~~~~~~~~~~~~~ CSVFile

// NOTE: This struct has the same format as TextFile and File.
// Convenient usage is: ReadContents() and WriteContents()
type CSVFile struct {
	*table.Table           // I want all Table methods to be available.
	file         *TextFile // I want to control access to file methods and overwrite them as need by
	hasHeadings  bool
}

// I would like if this fulfilled the RFC 4180 specs, however allowing multi line CSVs isn't a necessary feature for me currently.
// I will make it header agnostic, however
// RFC 4180 States:
// - headers are optional, system should work header agnostic
// - lines end with CRFL - I won't really bother with this unless there is a bug somewhere
// -

func NewCSVFile(filename string, headings bool) *CSVFile {
	filename = AddFileType(filename, "csv") // redundant check to see if the file is a CSV file, ensures it is if it isn't already
	file := NewTextFile(filename)

	table := table.NewTable(0)

	return &CSVFile{Table: table, file: file, hasHeadings: headings}
}

// returns an array of headings and a 2d array of
// Checks if there is a valid ab
func ReadCSV(filename string, headings bool) (csv *CSVFile, err error) {
	csv = NewCSVFile(filename, headings)
	err = csv.ReadContents()
	return csv, err
}

func (csv *CSVFile) String() string {
	return helpers.Format("File: %s\nHasHeadings: %t\nTable: \n%v\n", csv.file.Name(), csv.hasHeadings, csv.Table)
}

// Reading a CSV file
// Returns ErrInconsistentFieldNumber if the number of rows is inconsistent.
func (csv *CSVFile) ReadContents() (err error) {
	// Assumes the file attribute has been populated.
	fileContents := csv.file.ReadContents()

	csv.Table, err = csv.deserialise(fileContents)
	return err
}

// Converts a []string that is supposed to represent a CSV file's lines, to a table
func (csv *CSVFile) deserialise(contents []string) (*table.Table, error) {
	t := table.EmptyTable()
	err := errors.New("files/csv: Error occured while deserialising the given string array")
	var retErr error = nil
	errAdded := false

	// go through each row in contents
	// convert it to Row objects
	// add said Row objects to the table
	// return the table
	firstLine := true
	for line, content := range contents {
		// for each line:
		row := csvRecordToRow(content)

		if firstLine {
			firstLine = false
			t.Widen(row.Size())

			if csv.hasHeadings {
				_ = t.SetHeaders(row) // Error should not appear since the table has been widened to have the same size as the header row
				continue
			}

		}

		// TODO: To make compatible with RFC 4180, instead I could repeatedly cut a preffix using strings.Index(content, ",") and some more fancy logic

		// Add Record to table.
		recErr := t.AddRecord(row)
		if recErr != nil {
			if errors.Is(recErr, table.ErrIncompatibleSize) {
				// Convert to
				recErr = addErrInconsistentFieldNumber(line)
			}

			// Create new error, ErrInconsistentFieldNumber, reference the line number,
			lastLine := line == len(contents)-1

			if !lastLine {
				err = errors.Join(err, recErr)
				errAdded = true
			}
		}
	}

	if errAdded {
		retErr = err
	}
	return t, retErr
}

func csvRecordToRow(csvRecord string) *table.Row {
	cells := strings.Split(csvRecord, ",")
	row := table.NewRow(len(cells))

	for i, cell := range cells {
		_ = row.Set(i, cell) // Error ignored because the Row is set to have the same length as number of cells.
	}

	return row
}

func (c *CSVFile) WriteContents() {
	contents := c.serialise(c.Table) // TODO: handle this error

	c.file.ClearFile()
	c.file.AppendLines(contents, true)
	c.file.WriteContents()
}

// Converts from datastructure format (table), to TextFile format []string
func (csv *CSVFile) serialise(table *table.Table) (contents []string) {
	rows := table.Records()

	for _, row := range rows {
		line := row.Join("", "", ",", "", "")
		contents = append(contents, line)
	}

	return
}

func (c *CSVFile) Contents() []string {
	contents := c.serialise(c.Table)

	return contents
}

// TODO: Overwrite the below methods
// NOTE:
// All of the methods below should ONLY edit the csv.table parameter, and then when the user wants to finish writing,
// they should call WriteContents()

// // Returns -1 if: No headings, Heading not found
// // Otherwise: Returns the index of the heading.
// func (c *CSVFile) IndexOfCol(header string) (index int) {
// 	if !c.hasHeaders || len(c.contents) == 0 {
// 		return -1
// 	}
// 	headings := c.contents[0]
// 	for i, heading := range headings {
// 		if reflect.DeepEqual(heading, header) {
// 			index = i
// 		}
// 	}
//
// 	return
// }
//
// func (c *CSVFile) Row(i int) string { // TODO make more efficient
// 	return strings.Join(c.contents[i], ",")
// 	// return c.contentBuffer[i+1] // this is buggy, fix
// }
//
// func (c *CSVFile) Cell(row, col int) string {
// 	return c.contents[row][col]
// }
//
// func (c *CSVFile) Cols() int {
// 	if len(c.contents) == 0 {
// 		return 0
// 	} else {
// 		return len(c.contents[0])
// 	}
// }
//
// func (c *CSVFile) Rows() int {
// 	return len(c.contents)
// }
//
// func (csv *CSVFile) HasHeaders() bool {
// 	return csv.hasHeaders
// }
//
// func (csv *CSVFile) Headers() ([]string, error) {
// 	emptyHeaders := []string{}
//
// 	print(len(csv.contents))
// 	if csv.hasHeaders && len(csv.contents) > 0 {
// 		return csv.contents[0], nil
// 	}
//
// 	return emptyHeaders, ErrMissingHeaders
// }
