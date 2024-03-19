package errors

import "github.com/pkg/errors"

var (
	ErrHedgingMode = errors.New("Invalid Hedging Mode")
)
