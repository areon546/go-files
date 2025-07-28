package table

import (
	"errors"
)

// end of row signifies an index equal to the length, it implies index out of bounds
var ErrOutOfBounds = errors.New("files/table: index out of bounds")
