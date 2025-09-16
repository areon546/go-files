package files

import (
	"errors"
	"syscall"

	"github.com/areon546/go-helpers/helpers"
)

var (
	ErrPathEmpty         = errors.New("files: Path empty. ")
	errEmpty             = errors.New("")
	ErrNoFileOrDirectory = helpers.WrapError("files: %w", syscall.ENOENT) // special error for "no such file or directory" error that happens sometimes
)

func newErrNoFileOrDirectory(filename string) error {
	return helpers.WrapError("%w at file path: %s", ErrNoFileOrDirectory, filename)
}
