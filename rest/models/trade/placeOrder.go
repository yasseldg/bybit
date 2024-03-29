package trade

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sFloats"
)

type PlaceOrder struct {
	Order

	side             string
	orderType        string
	qty              string
	isLeverage       *int
	triggerDirection *int
	orderFilter      *string
	timeInForce      *string
	positionIdx      *int
	reduceOnly       *bool
	closeOnTrigger   *bool
	smpType          *string
	mmp              *bool
	tpOrderType      *string
	slOrderType      *string
}

func NewPlaceOrder(category, symbol, side, orderType string, qty float64, prec int) (*PlaceOrder, error) {
	order, err := newOrder(category, symbol, prec)
	if err != nil {
		return nil, err
	}

	if len(side) == 0 {
		return nil, errors.ErrInvalidSide
	}
	if len(orderType) == 0 {
		return nil, errors.ErrInvalidOrderType
	}
	if qty <= 0 {
		return nil, errors.ErrInvalidOrderQty
	}

	return &PlaceOrder{
		Order:     *order,
		side:      side,
		orderType: orderType,
		qty:       sFloats.ToString(qty, prec),
	}, nil
}

func (o *PlaceOrder) Log() {
	sLog.Info("PlaceOrder: %v ", o.GetParams())
}

func (o *PlaceOrder) IsLeverage(isLeverage int) *PlaceOrder {
	o.isLeverage = &isLeverage
	return o
}

func (o *PlaceOrder) TriggerDirection(direction int) *PlaceOrder {
	o.triggerDirection = &direction
	return o
}

func (o *PlaceOrder) OrderFilter(filter string) *PlaceOrder {
	o.orderFilter = &filter
	return o
}

func (o *PlaceOrder) TimeInForce(tif string) *PlaceOrder {
	o.timeInForce = &tif
	return o
}

func (o *PlaceOrder) PositionIdx(idx int) *PlaceOrder {
	o.positionIdx = &idx
	return o
}

func (o *PlaceOrder) ReduceOnly(reduce bool) *PlaceOrder {
	o.reduceOnly = &reduce
	return o
}

func (o *PlaceOrder) CloseOnTrigger(close bool) *PlaceOrder {
	o.closeOnTrigger = &close
	return o
}

func (o *PlaceOrder) SmpType(smp string) *PlaceOrder {
	o.smpType = &smp
	return o
}

func (o *PlaceOrder) Mmp(mmp bool) *PlaceOrder {
	o.mmp = &mmp
	return o
}

func (o *PlaceOrder) TpOrderType(orderType string) *PlaceOrder {
	o.tpOrderType = &orderType
	return o
}

func (o *PlaceOrder) SlOrderType(orderType string) *PlaceOrder {
	o.slOrderType = &orderType
	return o
}

// Params returns the order parameters
func (o *PlaceOrder) GetParams() common.Params {
	cp := o.Order.getParams()

	cp.Set("side", o.side)
	cp.Set("orderType", o.orderType)
	cp.Set("qty", o.qty)

	if o.isLeverage != nil {
		cp.Set("isLeverage", *o.isLeverage)
	}
	if o.triggerDirection != nil {
		cp.Set("triggerDirection", *o.triggerDirection)
	}
	if o.orderFilter != nil {
		cp.Set("orderFilter", *o.orderFilter)
	}
	if o.timeInForce != nil {
		cp.Set("timeInForce", *o.timeInForce)
	}
	if o.positionIdx != nil {
		cp.Set("positionIdx", *o.positionIdx)
	}
	if o.reduceOnly != nil {
		cp.Set("reduceOnly", *o.reduceOnly)
	}
	if o.closeOnTrigger != nil {
		cp.Set("closeOnTrigger", *o.closeOnTrigger)
	}
	if o.smpType != nil {
		cp.Set("smpType", *o.smpType)
	}
	if o.mmp != nil {
		cp.Set("mmp", *o.mmp)
	}
	if o.tpOrderType != nil {
		cp.Set("tpOrderType", *o.tpOrderType)
	}
	if o.slOrderType != nil {
		cp.Set("slOrderType", *o.slOrderType)
	}

	return cp
}
