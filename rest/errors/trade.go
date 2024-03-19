package errors

import "github.com/pkg/errors"

var (
	ErrInvalidOrderType = errors.New("Invalid order type")
	ErrInvalidOrderQty  = errors.New("Invalid order qty")
	ErrInvalidOrderId   = errors.New("Invalid order_id")
)
