package entity

import "errors"

var (
	ErrIDIsRequired            = errors.New("id is required")
	ErrInvalidID               = errors.New("invalid id")
	ErrClientIDIsRequired      = errors.New("client id is required")
	ErrInvalidClientID         = errors.New("invalid client id")
	ErrTransactionIDIsRequired = errors.New("transaction id is required")
	ErrInvalidTransactionID    = errors.New("invalid transaction id")
	ErrValueIsNegative         = errors.New("value must be positive")
	ErrValueIsZero             = errors.New("value must be greater than zero")
	ErrInvalidStatus           = errors.New("invalid status")
	ErrCannotReprocess         = errors.New("authorization cannot be reprocess, invalid status")
)
