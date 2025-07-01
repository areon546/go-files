package formatter

import "errors"

var (
	errEndOfRow = errors.New(": end of row")

	// end of row signifies an index equal to the length, it implies index out of bounds
	ErrOutOfBounds = errors.New("files/formatter: index out of bounds")
	ErrEndOfRow    = errors.Join(ErrOutOfBounds, errEndOfRow)
)
