package errors

import gerrors "errors"

var (
	ErrNotUnique      = gerrors.New("104")
	ErrValidation     = gerrors.New("103")
	ErrBadRequest     = gerrors.New("102")
	ErrRecordNotFound = gerrors.New("101")
	ErrUnauthorized   = gerrors.New("100")
)
