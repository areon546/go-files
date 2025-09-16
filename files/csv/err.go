package csv

import (
	"errors"

	"github.com/areon546/go-helpers/helpers"
)

var (
	errCSV = errors.New("csv: ")

	errMissingHeaders = errors.New("headings are missing")
	ErrMissingHeaders = helpers.WrapError("%w %w", errCSV, errMissingHeaders)

	errIncFieldNumber          = errors.New("number of fileds in CSV is inconstistent")
	ErrInconsistentFieldNumber = helpers.WrapError("%w %w", errCSV, errIncFieldNumber)
)

func newErrInconsistentFieldNumber(line int) error {
	return helpers.WrapError("%w at line %d", ErrInconsistentFieldNumber, line)
}
