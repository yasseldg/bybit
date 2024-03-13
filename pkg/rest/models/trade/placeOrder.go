package trade

import (
	"net/url"

	"github.com/yasseldg/bybit/pkg/rest/errors"

	"github.com/yasseldg/go-simple/types/sBool"
	"github.com/yasseldg/go-simple/types/sFloats"
	"github.com/yasseldg/go-simple/types/sInts"
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
func (o *PlaceOrder) GetParams() url.Values {
	p := o.Order.getParams()

	p.Add("side", o.side)
	p.Add("orderType", o.orderType)
	p.Add("qty", o.qty)

	if o.isLeverage != nil {
		p.Add("isLeverage", sInts.ToString(int64(*o.isLeverage)))
	}
	if o.triggerDirection != nil {
		p.Add("triggerDirection", sInts.ToString(int64(*o.triggerDirection)))
	}
	if o.orderFilter != nil {
		p.Add("orderFilter", *o.orderFilter)
	}
	if o.timeInForce != nil {
		p.Add("timeInForce", *o.timeInForce)
	}
	if o.positionIdx != nil {
		p.Add("positionIdx", sInts.ToString(int64(*o.positionIdx)))
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
	if o.tpOrderType != nil {
		p.Add("tpOrderType", *o.tpOrderType)
	}
	if o.slOrderType != nil {
		p.Add("slOrderType", *o.slOrderType)
	}

	return p
}
