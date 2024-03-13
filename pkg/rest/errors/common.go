package errors

import "github.com/pkg/errors"

var (
	ErrInvalidCategory  = errors.New("Invalid category")
	ErrInvalidSymbol    = errors.New("Invalid symbol")
	ErrInvalidSide      = errors.New("Invalid side")
	ErrInvalidPrecision = errors.New("Invalid precision")
)
