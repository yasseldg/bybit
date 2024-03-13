package trade

import (
	"net/url"

	"github.com/yasseldg/bybit/pkg/rest/errors"
)

type CancelOrder struct {
	category    string
	symbol      string
	orderId     string
	orderLinkId *string
	orderFilter *string
}

func NewCancelOrder(category, symbol, order_id string) (*CancelOrder, error) {
	if category == "" {
		return nil, errors.ErrInvalidCategory
	}
	if symbol == "" {
		return nil, errors.ErrInvalidSymbol
	}
	if order_id == "" {
		return nil, errors.ErrInvalidOrderId
	}

	return &CancelOrder{
		category: category,
		symbol:   symbol,
		orderId:  order_id,
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
func (o *CancelOrder) GetParams() url.Values {
	p := url.Values{}

	p.Add("category", o.category)
	p.Add("symbol", o.symbol)
	p.Add("orderId", o.orderId)

	if o.orderLinkId != nil {
		p.Add("orderLinkId", *o.orderLinkId)
	}
	if o.orderFilter != nil {
		p.Add("orderFilter", *o.orderFilter)
	}

	return p
}
