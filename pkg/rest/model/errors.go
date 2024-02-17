package model

import "github.com/pkg/errors"

var (
	errInvalidCategory  = errors.New("Invalid category")
	errInvalidSymbol    = errors.New("Invalid symbol")
	errInvalidSide      = errors.New("Invalid side")
	errInvalidOrderType = errors.New("Invalid order type")
	errInvalidOrderQty  = errors.New("Invalid order qty")
	errInvalidOrderId   = errors.New("Invalid order_id")
)
