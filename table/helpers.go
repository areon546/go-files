package table

import (
	"github.com/areon546/go-helpers/helpers"
)

func format(s string, a ...any) string { return helpers.Format(s, a...) }

func indexWithinBounds(i, max int) bool {
	return 0 <= i && i < max
}
