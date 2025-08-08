package files

import (
	"errors"

	"github.com/areon546/go-helpers/helpers"
)

var (
	ErrPathEmpty               = errors.New("files: Path empty. ")
	errEmpty                   = errors.New("")
	ErrInconsistentFieldNumber = errors.New("csv: Number of fields in CSV is inconsistent")
)

func addErrInconsistentFieldNumber(line int) error {
	return helpers.WrapError("%w at line %d", ErrInconsistentFieldNumber, line)
}
