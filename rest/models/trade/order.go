package trade

import (
	"fmt"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"

	"github.com/yasseldg/go-simple/types/sFloats"
)

type Order struct {
	category     string
	symbol       string
	prec         int
	price        *string
	triggerPrice *string
	triggerBy    *string
	orderIv      *string
	orderLinkId  *string
	tpslMode     *string
	takeProfit   *string
	stopLoss     *string
	tpTriggerBy  *string
	slTriggerBy  *string
	tpLimitPrice *string
	slLimitPrice *string
}

func newOrder(category, symbol string, prec int) (*Order, error) {
	if len(category) == 0 {
		return nil, errors.ErrInvalidCategory
	}

	if len(symbol) == 0 {
		return nil, errors.ErrInvalidSymbol
	}

	if prec <= 0 {
		return nil, errors.ErrInvalidPrecision
	}

	return &Order{
		category: category,
		symbol:   symbol,
		prec:     prec,
	}, nil
}

func (o *Order) String() string {
	return fmt.Sprintf("%v", o.getParams())
}

func (o *Order) TriggerPrice(triggerPrice float64) *Order {
	trp := sFloats.ToString(triggerPrice, o.prec)
	o.triggerPrice = &trp
	return o
}

func (o *Order) OrderLinkId(orderLinkId string) *Order {
	o.orderLinkId = &orderLinkId
	return o
}

func (o *Order) Price(price float64) *Order {
	p := sFloats.ToString(price, o.prec)
	o.price = &p
	return o
}

func (o *Order) TriggerBy(triggerBy string) *Order {
	o.triggerBy = &triggerBy
	return o
}

func (o *Order) OrderIv(iv string) *Order {
	o.orderIv = &iv
	return o
}

func (o *Order) TakeProfit(profit float64) *Order {
	tp := sFloats.ToString(profit, o.prec)
	o.takeProfit = &tp
	return o
}

func (o *Order) StopLoss(loss float64) *Order {
	sl := sFloats.ToString(loss, o.prec)
	o.stopLoss = &sl
	return o
}

func (o *Order) TpTriggerBy(triggerBy string) *Order {
	o.tpTriggerBy = &triggerBy
	return o
}

func (o *Order) SlTriggerBy(triggerBy string) *Order {
	o.slTriggerBy = &triggerBy
	return o
}

func (o *Order) TpslMode(mode string) *Order {
	o.tpslMode = &mode
	return o
}

func (o *Order) TpLimitPrice(price float64) *Order {
	p := sFloats.ToString(price, o.prec)
	o.tpLimitPrice = &p
	return o
}

func (o *Order) SlLimitPrice(price float64) *Order {
	p := sFloats.ToString(price, o.prec)
	o.slLimitPrice = &p
	return o
}

// Params returns the order parameters
func (o *Order) getParams() common.Params {
	cp := common.NewParams()

	cp.Set("category", o.category)
	cp.Set("symbol", o.symbol)

	if o.price != nil {
		cp.Set("price", *o.price)
	}

	if o.triggerPrice != nil {
		cp.Set("triggerPrice", *o.triggerPrice)
	}
	if o.triggerBy != nil {
		cp.Set("triggerBy", *o.triggerBy)
	}
	if o.orderIv != nil {
		cp.Set("orderIv", *o.orderIv)
	}
	if o.orderLinkId != nil {
		cp.Set("orderLinkId", *o.orderLinkId)
	}
	if o.takeProfit != nil {
		cp.Set("takeProfit", *o.takeProfit)
	}
	if o.stopLoss != nil {
		cp.Set("stopLoss", *o.stopLoss)
	}
	if o.tpTriggerBy != nil {
		cp.Set("tpTriggerBy", *o.tpTriggerBy)
	}
	if o.slTriggerBy != nil {
		cp.Set("slTriggerBy", *o.slTriggerBy)
	}
	if o.tpslMode != nil {
		cp.Set("tpslMode", *o.tpslMode)
	}
	if o.tpLimitPrice != nil {
		cp.Set("tpLimitPrice", *o.tpLimitPrice)
	}
	if o.slLimitPrice != nil {
		cp.Set("slLimitPrice", *o.slLimitPrice)
	}

	return cp
}
