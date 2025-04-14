package domain

import "errors"

var (
	ErrAccountNotFound      = errors.New("account not found")
	ErrDuplicateApiKey      = errors.New("duplicate api key")
	ErrInvoiceNotFound      = errors.New("invoice not found")
	ErrUnauthorizedAccess   = errors.New("unauthorized access")
	ErrAccountAlreadyExists = errors.New("account already exists")
)
