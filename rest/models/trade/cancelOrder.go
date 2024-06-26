package trade

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"
)

type CancelOrder struct {
	category    string
	symbol      string
	orderId     string
	orderLinkId *string
	orderFilter *string
}

func NewCancelOrder(category, symbol, orderId string) (*CancelOrder, error) {
	if category == "" {
		return nil, errors.ErrInvalidCategory
	}
	if symbol == "" {
		return nil, errors.ErrInvalidSymbol
	}
	if orderId == "" {
		return nil, errors.ErrInvalidOrderId
	}

	return &CancelOrder{
		category: category,
		symbol:   symbol,
		orderId:  orderId,
	}, nil
}

func (o *CancelOrder) OrderLinkId(orderLinkId string) *CancelOrder {
	o.orderLinkId = &orderLinkId
	return o
}

func (o *CancelOrder) OrderFilter(filter string) *CancelOrder {
	o.orderFilter = &filter
	return o
}

// Params returns the order parameters
func (o *CancelOrder) GetParams() common.Params {
	cp := common.NewParams()

	cp.Set("category", o.category)
	cp.Set("symbol", o.symbol)
	cp.Set("orderId", o.orderId)

	if o.orderLinkId != nil {
		cp.Set("orderLinkId", *o.orderLinkId)
	}
	if o.orderFilter != nil {
		cp.Set("orderFilter", *o.orderFilter)
	}

	return cp
}
