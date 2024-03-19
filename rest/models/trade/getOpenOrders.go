package trade

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"
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
func (o *GetOpenOrders) GetParams() common.Params {
	cp := common.NewParams()

	cp.Set("category", o.category)

	if o.symbol != nil {
		cp.Set("symbol", *o.symbol)
	}
	if o.baseCoin != nil {
		cp.Set("baseCoin", *o.baseCoin)
	}
	if o.settleCoin != nil {
		cp.Set("settleCoin", *o.settleCoin)
	}
	if o.orderId != nil {
		cp.Set("orderId", *o.orderId)
	}
	if o.orderLinkId != nil {
		cp.Set("orderLinkId", *o.orderLinkId)
	}
	if o.orderFilter != nil {
		cp.Set("orderFilter", *o.orderFilter)
	}

	return cp
}
