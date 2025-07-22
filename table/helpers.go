package table

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

func handle(err error) {
	// helpers.Handle(err)
	helpers.Handle(err)
}

func print(a ...any) {
	helpers.Print(a...)
}

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func assertEqualsObject(t testing.TB, expected, result any) {
	t.Helper()
	helpers.AssertEqualsObject(t, expected, result)
}

func handleErrorExcept(t testing.TB, err, except error) {
	t.Helper()
	helpers.HandleExcept(err, except)
}
