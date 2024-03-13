package trade

import (
	"net/url"

	"github.com/yasseldg/bybit/pkg/rest/errors"
)

type GetOpenOrders struct {
	category    string
	symbol      *string
	baseCoin    *string
	settleCoin  *string
	orderId     *string
	orderLinkId *string
	orderFilter *string
}

func NewGetOpenOrders(category string) (*GetOpenOrders, error) {
	if len(category) == 0 {
		return nil, errors.ErrInvalidCategory
	}

	return &GetOpenOrders{
		category: category,
	}, nil
}

func (o *GetOpenOrders) Symbol(symbol string) *GetOpenOrders {
	o.symbol = &symbol
	return o
}

func (o *GetOpenOrders) BaseCoin(baseCoin string) *GetOpenOrders {
	o.baseCoin = &baseCoin
	return o
}

func (o *GetOpenOrders) SettleCoin(settleCoin string) *GetOpenOrders {
	o.settleCoin = &settleCoin
	return o
}

func (o *GetOpenOrders) OrderId(id string) *GetOpenOrders {
	o.orderId = &id
	return o
}

func (o *GetOpenOrders) OrderLinkId(orderLinkId string) *GetOpenOrders {
	o.orderLinkId = &orderLinkId
	return o
}

func (o *GetOpenOrders) OrderFilter(filter string) *GetOpenOrders {
	o.orderFilter = &filter
	return o
}

// Params returns the order parameters
func (o *GetOpenOrders) GetParams() url.Values {
	p := url.Values{}

	p.Add("category", o.category)

	if o.symbol != nil {
		p.Add("symbol", *o.symbol)
	}
	if o.baseCoin != nil {
		p.Add("baseCoin", *o.baseCoin)
	}
	if o.settleCoin != nil {
		p.Add("settleCoin", *o.settleCoin)
	}
	if o.orderId != nil {
		p.Add("orderId", *o.orderId)
	}
	if o.orderLinkId != nil {
		p.Add("orderLinkId", *o.orderLinkId)
	}
	if o.orderFilter != nil {
		p.Add("orderFilter", *o.orderFilter)
	}

	return p
}
