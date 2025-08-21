package files

import (
	"errors"
	"syscall"

	"github.com/areon546/go-helpers/helpers"
)

var (
	ErrPathEmpty               = errors.New("files: Path empty. ")
	errEmpty                   = errors.New("")
	ErrInconsistentFieldNumber = errors.New("csv: Number of fields in CSV is inconsistent")
	ErrNoFileOrDirectory       = helpers.WrapError("files: %w", syscall.ENOENT) // special error for "no such file or directory" error that happens sometimes
)

func newErrInconsistentFieldNumber(line int) error {
	return helpers.WrapError("%w at line %d", ErrInconsistentFieldNumber, line)
}

func newErrNoFileOrDirectory(filename string) error {
	return helpers.WrapError("%w at file path: %s", ErrNoFileOrDirectory, filename)
}
