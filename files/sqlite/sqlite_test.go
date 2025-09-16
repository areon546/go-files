package sqlite

import (
	"fmt"
	"testing"

	"github.com/areon546/go-files/files"
	"github.com/areon546/go-helpers/helpers"
)

func Test(t *testing.T) {
	files.RemoveAllWithinDirectory("./test")

	fn := "file:./test/sqlite.db"

	n, _ := Open(fn)

	fmt.Println(n)
	res, err := n.Exec("CREATE TABLE IF NOT EXISTS test (name);")
	helpers.AssertNoError(t, err)

	rows, err := res.RowsAffected()
	helpers.AssertNoError(t, err)
	helpers.AssertEqualsObject(t, int64(0), rows)

	res, err = n.Exec("INSERT INTO test VALUES ('LOVE');")
	helpers.AssertNoError(t, err)

	rows, err = res.RowsAffected()
	helpers.AssertNoError(t, err)
	helpers.AssertEqualsObject(t, int64(1), rows)
}
