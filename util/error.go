package util

import "errors"

var (
	ErrSystem                   = errors.New("SystemError")
	ErrToken                    = errors.New("TokenError")
	ErrScope                    = errors.New("ScopeError")
	ErrNotFound                 = errors.New("NotFound")
	ErrMethodNotFound           = errors.New("MethodNotFound")
	ErrServiceNotFound          = errors.New("ServiceNotFound")
	ErrFileOrDirectoryNotExists = errors.New("FileOrDirectoryNotExists")
)
