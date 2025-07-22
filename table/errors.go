package table

import (
	"errors"
	"fmt"
)

var (
	errEndOfRow = errors.New(": end of row")

	// end of row signifies an index equal to the length, it implies index out of bounds
	ErrOutOfBounds = errors.New("files/table: index out of bounds")
	ErrEndOfRow    = fmt.Errorf("%w%w", ErrOutOfBounds, errEndOfRow)
)
