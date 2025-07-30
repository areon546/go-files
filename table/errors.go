package table

import (
	"errors"
)

// end of row signifies an index equal to the length, it implies index out of bounds
var (
	ErrHeaderMissing    error = errors.New("table: header missing from table")
	ErrOutOfBounds            = errors.New("table: index out of bounds")
	ErrCellPopulated          = errors.New("table: cell already has a value, overwritting")
	ErrIncompatibleSize       = errors.New("table: the values you are trying to insert into the table have a different size from what was expected")
)
