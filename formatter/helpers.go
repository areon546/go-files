package formatter

import (
	"testing"

	"github.com/areon546/go-helpers/helpers"
)

// helper functions

func handle(err error) {
	// helpers.Handle(err)
	helpers.Handle(err)
}

func print(a ...any) {
	helpers.Print(a...)
}

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func assertEqualsObject(t testing.TB, expected, result any) {
	helpers.AssertEqualsObject(t, expected, result)
}
