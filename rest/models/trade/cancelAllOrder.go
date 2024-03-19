package trade

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"
)

type CancelAllOrders struct {
	category      string
	symbol        *string
	baseCoin      *string
	settleCoin    *string
	orderFilter   *string
	stopOrderType *string
}

func NewCancelAllOrders(category string) (*CancelAllOrders, error) {
	if category == "" {
		return nil, errors.ErrInvalidCategory
	}

	return &CancelAllOrders{
		category: category,
	}, nil
}

func (o *CancelAllOrders) Symbol(symbol string) *CancelAllOrders {
	o.symbol = &symbol
	return o
}

func (o *CancelAllOrders) BaseCoin(baseCoin string) *CancelAllOrders {
	o.baseCoin = &baseCoin
	return o
}

func (o *CancelAllOrders) SettleCoin(settleCoin string) *CancelAllOrders {
	o.settleCoin = &settleCoin
	return o
}

func (o *CancelAllOrders) OrderFilter(orderFilter string) *CancelAllOrders {
	o.orderFilter = &orderFilter
	return o
}

func (o *CancelAllOrders) StopOrderType(stopOrderType string) *CancelAllOrders {
	o.stopOrderType = &stopOrderType
	return o
}

// Params returns the order parameters
func (o *CancelAllOrders) GetParams() common.Params {
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

	if o.orderFilter != nil {
		cp.Set("orderFilter", *o.orderFilter)
	}

	if o.stopOrderType != nil {
		cp.Set("stopOrderType", *o.stopOrderType)
	}

	return cp
}

// category	true	string	Product type

//     Unified account: spot, linear, inverse, option
//     Classic account: spot, linear, inverse

// symbol	false	string	Symbol name. linear & inverse: Required if not passing baseCoin or settleCoin
// baseCoin	false	string	Base coin

//     linear & inverse(Classic account): If cancel all by baseCoin, it will cancel all linear & inverse orders. Required if not passing symbol or settleCoin
//     linear & inverse(Unified account): If cancel all by baseCoin, it will cancel all corresponding category orders. Required if not passing symbol or settleCoin
//     Classic spot: invalid

// settleCoin	false	string	Settle coin

//     linear & inverse: Required if not passing symbol or baseCoin
//     Does not support spot

// orderFilter	false	string

//     category=spot, you can pass Order, tpslOrder, StopOrder, OcoOrder, BidirectionalTpslOrder. If not passed, Order by default
//     category=linear or inverse, you can pass Order, StopOrder. If not passed, all kinds of orders will be cancelled, like active order, conditional order, TP/SL order and trailing stop order
//     category=option, you can pass Order. No matter it is passed or not, always cancel all orders

// stopOrderType	false	string	Stop order type Stop
// Only used for category=linear or inverse and orderFilter=StopOrder,you can cancel conditional orders except TP/SL order and Trailing stop orders with this param
