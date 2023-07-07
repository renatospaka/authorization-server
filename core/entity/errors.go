package entity

import "errors"

var (
	ErrIDIsRequired         = errors.New("id is required")
	ErrInvalidID            = errors.New("invalid id")
	ErrClientIDIsRequired   = errors.New("client id is required")
	ErrInvalidClientID      = errors.New("invalid client id")
	ErrInvalidTransactionID = errors.New("invalid transaction id")
	ErrValueIsNegative      = errors.New("value must be positive")
	ErrValueIsZero          = errors.New("value must be greater than zero")
	ErrInvalidStatus        = errors.New("invalid statuso")
)
