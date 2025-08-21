package files

import (
	"testing"

	"github.com/areon546/go-files/table"
	"github.com/areon546/go-helpers/helpers"
)

var (
	standardCSV *CSVFile
	errFieldCSV *CSVFile
	headersCSV  *CSVFile
)

func init() {
	print("Initialising CSV_TEST")
	defer print("Finished Initialising CSV_TEST\n")

	standardCSV = NewCSVFile("./files/abc.csv", false)
	errFieldCSV = NewCSVFile("./files/arbitraryFields.csv", false)
	headersCSV = NewCSVFile("./files/headers.csv", true)
}

func TestReadContents(t *testing.T) {
	// Returns an error with wrong types of

	t.Run("No Errors Expected", func(t *testing.T) {
		csv := standardCSV

		tab, err := csv.ReadContents()
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsObject(t, csv.Table, tab)

		helpers.AssertEqualsInt(t, 2, csv.Entries())
		helpers.AssertEqualsInt(t, 4, csv.Width())
	})

	t.Run("Headers", func(t *testing.T) {
		csv := headersCSV

		tab, err := csv.ReadContents()
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsObject(t, csv.Table, tab)

		head, err := csv.Headers()
		helpers.AssertNoError(t, err)
		helpers.AssertEqualsInt(t, 3, head.Size())

		header1, err := head.Get(0)
		helpers.AssertNoError(t, err)
		helpers.AssertEquals(t, "skinName", header1)
	})

	t.Run("ErrInconsistentFieldNumber Expected", func(t *testing.T) {
		csv := errFieldCSV

		expectedTab := table.NewTable(2)
		r1 := table.NewRow(2)
		r1.Set(0, "sd")
		r2 := table.NewRow(2)
		r2.Set(0, "as")
		r2.Set(1, "as")
		expectedTab.AddRecord(r1)
		expectedTab.AddRecord(r2)

		tab, err := csv.ReadContents()
		helpers.AssertError(t, err, ErrInconsistentFieldNumber)

		helpers.AssertEqualsObject(t, expectedTab, tab)
	})

	// helpers.AssertEqualsInt(t, 0, 1)
}

func TestCSVWriteContents(t *testing.T) {
	t.Run("add record", func(t *testing.T) {
		tab := standardCSV
		helpers.AssertEqualsInt(t, 2, tab.Entries())

		r := table.NewRow(tab.Width())
		r.Set(0, "4")
		r.Set(1, "8")
		r.Set(2, "1")
		r.Set(3, "2")

		tab.AddRecord(r)
		helpers.AssertEqualsInt(t, 3, tab.Entries())
	})
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
