package pgxpoolmock

import (
	"errors"
)

var (
	ErrToBeCalledWithPointers             = errors.New("expected scan to be called with pointers")
	ErrUnexpectedNilVal                   = errors.New("unexpected nil value for arg")
	ErrScanExpectedToHaveSameNumberOfArgs = errors.New("expected scan to be called with same number of arguments")
	ErrCantSetDestination                 = errors.New("cannot set destination value")
	ErrNotSupported                       = errors.New("not supported")
	ErrMustBeAPointer                     = errors.New("destination argument must be a pointer")
	ErrIncorrectArgNumber                 = errors.New("incorrect argument number")
	ErrUnexpectedArg                      = errors.New("scan with unexpected arg")
	ErrEndBatchResult                     = errors.New("batch already closed") // Use the error to signify the end of a batch result.
	ErrNoBatchResult                      = errors.New("no result")
)
