package files

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

var (
	standardCSV *CSVFile
	errFieldCSV *CSVFile
	headersCSV  *CSVFile
)

func init() {
	print("Initialising CSV_TEST")
	defer print("Finished Initialising CSV_TEST")

	standardCSV = NewCSVFile("./files/abc.csv", false)
	errFieldCSV = NewCSVFile("./files/arbitraryFields.csv", false)
	headersCSV = NewCSVFile("./files/headers.csv", true)
}

func TestReadContents(t *testing.T) {
	// Returns an error with wrong types of

	t.Run("No Errors Expected", func(t *testing.T) {
		csv := standardCSV

		err := csv.ReadContents()
		helpers.AssertNoError(t, err)

		helpers.Print("Width", csv.table.Width())
		helpers.Print(csv)
		helpers.AssertEqualsInt(t, 2, csv.table.Entries())
	})

	t.Run("Headers", func(t *testing.T) {
		csv := headersCSV
		err := csv.ReadContents()

		helpers.AssertNoError(t, err)
	})

	t.Run("ErrInconsistentFieldNumber Expected", func(t *testing.T) {
		csv := errFieldCSV
		err := csv.ReadContents()

		helpers.Print(csv)

		// helpers.AssertEqualsInt(t, 3, csv.Rows())
		helpers.AssertError(t, err, ErrInconsistentFieldNumber)
	})

	// helpers.AssertEqualsInt(t, 0, 1)
}

//
// func TestHeaders(t *testing.T) {
// 	t.Run("Expect to have headers", func(t *testing.T) {
// 		csv := headersCSV
// 		err := csv.ReadContents()
//
// 		helpers.AssertNoError(t, err)
//
// 		headers, err := csv.Headers()
//
// 		helpers.AssertNoError(t, err)
// 		helpers.AssertEqualsObject(t, []string{"skinName", "address", "id"}, headers)
// 	})
// 	// check for errors:
// 	// ErrHeadernotFound
//
// 	t.Run("Test Missing Headers", func(t *testing.T) {
// 		csv := standardCSV
// 		err := csv.ReadContents()
//
// 		helpers.AssertNoError(t, err)
//
// 		headers, err := csv.Headers()
//
// 		helpers.AssertError(t, err, ErrMissingHeaders)
// 		helpers.AssertEqualsObject(t, []string{}, headers)
// 	})
// }
