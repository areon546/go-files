package sqlite

import (
	"fmt"
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func Test(t *testing.T) {
	fn := "file:sqlite.db"

	n, _ := Open(fn)

	fmt.Println(n)
	n.Query("CREATE TABLE IF NOT EXISTS skins (id, name);")

	helpers.AssertEquals(t, "", "a2")
}
