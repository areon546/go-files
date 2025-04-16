package files

import (
	"reflect"
	"strings"
)

// ~~~~~~~~~~~~~~~~~~~~ CSVFile
type CSVFile struct { // TODO convert to table datastructure
	TextFile
	headings []string
	contents [][]string
}

// returns an array of headings and a 2d array of
func ReadCSV(filename string) (csv CSVFile) {

	// check if fileName ends with csv
	filename = trimFiletype(filename, "csv")

	file := NewTextFile(ConstructFilePath("", filename, "csv"))
	// read fileName into CSVFile

	// file := makeFile(fileName)
	fileContents := file.ReadFile()

	// go through each line in CSV and
	for i, csvCell := range fileContents {
		// print("csv:", csvCell)
		if i == 0 { // adds headings to headings attribute
			csv.headings = strings.Split(csvCell, ",")
		} else { // ads csv items to contents attribute

			// check if the string is empty, if so skip
			if reflect.DeepEqual(csvCell, "") {
				continue
			}

			csv.contents = append(csv.contents, strings.Split(csvCell, ","))
		}
	}

	return
}

func (c *CSVFile) IndexOfCol(header string) (index int) {
	for i, heading := range c.headings {
		if reflect.DeepEqual(heading, header) {
			index = i
		}
	}

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
	return len(c.headings)
}

func (c *CSVFile) Rows() int {
	return len(c.contents)
}

func (c *CSVFile) Contents() [][]string {
	return c.contents
}
