package model

import (
	"net/url"

	"github.com/yasseldg/go-simple/types/sBool"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
)

type Order struct {
	category         string
	symbol           string
	isLeverage       *int
	side             string
	orderType        string
	qty              string
	price            *string
	triggerDirection *int
	orderFilter      *string
	triggerPrice     *string
	triggerBy        *string
	orderIv          *string
	timeInForce      *string
	positionIdx      *int
	orderLinkId      *string
	takeProfit       *string
	stopLoss         *string
	tpTriggerBy      *string
	slTriggerBy      *string
	reduceOnly       *bool
	closeOnTrigger   *bool
	smpType          *string
	mmp              *bool
	tpslMode         *string
	tpLimitPrice     *string
	slLimitPrice     *string
	tpOrderType      *string
	slOrderType      *string
}

func NewOrder(category, symbol, side, orderType string, qty float64) (*Order, error) {
	if len(category) == 0 {
		return nil, errInvalidCategory
	}

	if len(symbol) == 0 {
		return nil, errInvalidSymbol
	}

	if len(side) == 0 {
		return nil, errInvalidSide
	}

	if len(orderType) == 0 {
		return nil, errInvalidOrderType
	}

	if qty <= 0 {
		return nil, errInvalidOrderQty
	}

	return &Order{
		category:  category,
		symbol:    symbol,
		side:      side,
		orderType: orderType,
		qty:       sFloats.ToString(qty),
	}, nil
}

func (o *Order) TimeInForce(tif string) *Order {
	o.timeInForce = &tif
	return o
}

func (o *Order) IsLeverage(isLeverage int) *Order {
	o.isLeverage = &isLeverage
	return o
}

func (o *Order) TriggerPrice(triggerPrice string) *Order {
	o.triggerPrice = &triggerPrice
	return o
}

func (o *Order) OrderLinkId(orderLinkId string) *Order {
	o.orderLinkId = &orderLinkId
	return o
}

func (o *Order) Price(price string) *Order {
	o.price = &price
	return o
}

func (o *Order) TriggerDirection(direction int) *Order {
	o.triggerDirection = &direction
	return o
}

func (o *Order) OrderFilter(filter string) *Order {
	o.orderFilter = &filter
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

func (o *Order) PositionIdx(idx int) *Order {
	o.positionIdx = &idx
	return o
}

func (o *Order) TakeProfit(profit string) *Order {
	o.takeProfit = &profit
	return o
}

func (o *Order) StopLoss(loss string) *Order {
	o.stopLoss = &loss
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

func (o *Order) ReduceOnly(reduce bool) *Order {
	o.reduceOnly = &reduce
	return o
}

func (o *Order) CloseOnTrigger(close bool) *Order {
	o.closeOnTrigger = &close
	return o
}

func (o *Order) SmpType(smp string) *Order {
	o.smpType = &smp
	return o
}

func (o *Order) Mmp(mmp bool) *Order {
	o.mmp = &mmp
	return o
}

func (o *Order) TpslMode(mode string) *Order {
	o.tpslMode = &mode
	return o
}

func (o *Order) TpLimitPrice(price string) *Order {
	o.tpLimitPrice = &price
	return o
}

func (o *Order) SlLimitPrice(price string) *Order {
	o.slLimitPrice = &price
	return o
}

func (o *Order) TpOrderType(orderType string) *Order {
	o.tpOrderType = &orderType
	return o
}

func (o *Order) SlOrderType(orderType string) *Order {
	o.slOrderType = &orderType
	return o
}

// Params returns the order parameters
func (o *Order) GetParams() url.Values {
	p := url.Values{}

	p.Add("category", o.category)
	p.Add("symbol", o.symbol)
	p.Add("side", o.side)
	p.Add("orderType", o.orderType)
	p.Add("qty", o.qty)

	if o.price != nil {
		p.Add("price", *o.price)
	}
	if o.triggerDirection != nil {
		p.Add("triggerDirection", sInts.ToString(int64(*o.triggerDirection)))
	}
	if o.orderFilter != nil {
		p.Add("orderFilter", *o.orderFilter)
	}
	if o.triggerPrice != nil {
		p.Add("triggerPrice", *o.triggerPrice)
	}
	if o.triggerBy != nil {
		p.Add("triggerBy", *o.triggerBy)
	}
	if o.orderIv != nil {
		p.Add("orderIv", *o.orderIv)
	}
	if o.timeInForce != nil {
		p.Add("timeInForce", *o.timeInForce)
	}
	if o.positionIdx != nil {
		p.Add("positionIdx", sInts.ToString(int64(*o.positionIdx)))
	}
	if o.orderLinkId != nil {
		p.Add("orderLinkId", *o.orderLinkId)
	}
	if o.takeProfit != nil {
		p.Add("takeProfit", *o.takeProfit)
	}
	if o.stopLoss != nil {
		p.Add("stopLoss", *o.stopLoss)
	}
	if o.tpTriggerBy != nil {
		p.Add("tpTriggerBy", *o.tpTriggerBy)
	}
	if o.slTriggerBy != nil {
		p.Add("slTriggerBy", *o.slTriggerBy)
	}
	if o.reduceOnly != nil {
		p.Add("reduceOnly", sBool.ToString(*o.reduceOnly))
	}
	if o.closeOnTrigger != nil {
		p.Add("closeOnTrigger", sBool.ToString(*o.closeOnTrigger))
	}
	if o.smpType != nil {
		p.Add("smpType", *o.smpType)
	}
	if o.mmp != nil {
		p.Add("mmp", sBool.ToString(*o.mmp))
	}
	if o.tpslMode != nil {
		p.Add("tpslMode", *o.tpslMode)
	}
	if o.tpLimitPrice != nil {
		p.Add("tpLimitPrice", *o.tpLimitPrice)
	}
	if o.slLimitPrice != nil {
		p.Add("slLimitPrice", *o.slLimitPrice)
	}
	if o.tpOrderType != nil {
		p.Add("tpOrderType", *o.tpOrderType)
	}
	if o.slOrderType != nil {
		p.Add("slOrderType", *o.slOrderType)
	}

	return p
}
