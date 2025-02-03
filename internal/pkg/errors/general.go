package errors

import "errors"

var (
	ErrInvalidRequestParameter = errors.New("invalid request parameter")
	ErrInternalServerError     = errors.New("internal system error")
)
