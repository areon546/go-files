package files

import (
	"errors"
	"reflect"
	"strings"

	"github.com/areon546/go-helpers/helpers"
)

var (
	ErrInconsistentFieldNumber error = errors.New("CSV: Number of fields in CSV is inconsistent")
	errMissingHeaders          error = errors.New("Headings are missing")
	ErrMissingHeaders          error = helpers.WrapError("%w %w", errFiles, errMissingHeaders)
)

// ~~~~~~~~~~~~~~~~~~~~ CSVFile
type CSVFile struct {
	// TODO: ? convert to table datastructure
	file     TextFile
	contents [][]string

	hasHeaders bool
	width      int

	// TODO: Add restrictions for using TextFile functionality, I don't want the user to be able to miscalaneously write XYZ without it updating the CSV datastructure.
}

// I would like if this fulfilled the RFC 4180 specs, however allowing multi line CSVs isn't a necessary feature for me currently.
// I will make it header agnostic, however
// RFC 4180 States:
// - headers are optional, system should work header agnostic
// - lines end with CRFL - I won't really bother with this unless there is a bug somewhere
// -

func NewCSVFile(filename string, headings bool) *CSVFile {
	filename = trimFiletype(filename, "csv")
	hasCSVSuff := strings.HasSuffix(filename, ".csv")
	filePath := filename
	if !hasCSVSuff {
		filePath += ".csv"
	}
	file := NewTextFile(filePath)

	return &CSVFile{file: *file, hasHeaders: headings, contents: [][]string{}}
}

// returns an array of headings and a 2d array of
// Checks if there is a valid ab
func ReadCSV(filename string, headings bool) (csv *CSVFile, err error) {
	csv = NewCSVFile(filename, headings)
	err = csv.ReadContents()
	return csv, err
}

func (c *CSVFile) IndexOfCol(header string) (index int) {
	// for i, heading := range c.headings {
	// 	if reflect.DeepEqual(heading, header) {
	// 		index = i
	// 	}
	// }

	return
}

func (c *CSVFile) Row(i int) string { // TODO make more efficient
	return strings.Join(c.contents[i], ",")
	// return c.contentBuffer[i+1] // this is buggy, fix
}

func (c *CSVFile) Cell(row, col int) string {
	return c.contents[row][col]
}

func (c *CSVFile) Cols() int {
	if len(c.contents) == 0 {
		return 0
	} else {
		return len(c.contents[0])
	}
}

func (c *CSVFile) Rows() int {
	return len(c.contents)
}

func (csv *CSVFile) HasHeaders() bool {
	return csv.hasHeaders
}

func (csv *CSVFile) Headers() ([]string, error) {
	emptyHeaders := []string{}

	print(len(csv.contents))
	if csv.hasHeaders && len(csv.contents) > 0 {
		return csv.contents[0], nil
	}

	return emptyHeaders, ErrMissingHeaders
}

// Reading a CSV file
// Returns ErrInconsistentFieldNumber if the number of rows is inconsistent.
func (csv *CSVFile) ReadContents() (err error) {
	err = errEmpty

	contents := csv.file.ReadFile()
	firstLine := true
	var csvContent [][]string

	// Go through each line in the TextFile
	for index, record := range contents {
		line := index + 1
		// Go to the next line, if the line in the record is empty.
		if reflect.DeepEqual(record, "") {
			continue
		} else {
			cells := strings.Split(record, ",")

			if firstLine {
				// first line indicates headings,
				csv.width = len(cells)
				firstLine = false
			}

			// Check for errors
			if csv.width != len(cells) {
				lineError := helpers.WrapError("%w at line: %d", ErrInconsistentFieldNumber, line)
				err = errors.Join(err, lineError) // Say: Happens on index XYZ
			}

			csvContent = append(csvContent, cells)

			// TODO: Honestly, for now it would take a lot of effort to make this coincide with the  RFC 4180
			// specification of a CSV file, that for now I will accept it not being completely inline with the specs. I will simply update relevant comments.
		}
	}
	csv.contents = csvContent

	return returnErr(err)
}

/* Overwritten Functions */

func (c *CSVFile) Contents() [][]string {
	return c.contents
}
