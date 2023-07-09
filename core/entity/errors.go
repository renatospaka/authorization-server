package entity

import "errors"

var (
	ErrAuthorizationIDIsRequired = errors.New("authorization id is required")
	ErrInvalidAuthorizationID    = errors.New("invalid authorization id")
	ErrAuthorizationIDNotFound   = errors.New("authorization not found")
	ErrClientIDIsRequired        = errors.New("client id is required")
	ErrInvalidClientID           = errors.New("invalid client id")
	ErrTransactionIDIsRequired   = errors.New("transaction id is required")
	ErrInvalidTransactionID      = errors.New("invalid transaction id")
	ErrValueIsNegative           = errors.New("value must be positive")
	ErrValueIsZero               = errors.New("value must be greater than zero")
	ErrInvalidStatus             = errors.New("invalid status")
	ErrCannotReprocess           = errors.New("authorization cannot be reprocessed, status not allowed")
)
