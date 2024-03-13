package trade

import (
	"net/url"

	"github.com/yasseldg/bybit/pkg/rest/errors"

	"github.com/yasseldg/go-simple/types/sFloats"
)

type AmendOrder struct {
	Order

	orderId string
	qty     *string
}

func NewAmendOrder(category, symbol, orderId string, prec int) (*AmendOrder, error) {
	order, err := newOrder(category, symbol, prec)
	if err != nil {
		return nil, err
	}

	if len(orderId) == 0 {
		return nil, errors.ErrInvalidOrderId
	}

	return &AmendOrder{
		Order:   *order,
		orderId: orderId,
	}, nil
}

func (o *AmendOrder) Qty(qty float64) *AmendOrder {
	q := sFloats.ToString(qty, o.prec)
	o.qty = &q
	return o
}

// Params returns the order parameters
func (o *AmendOrder) GetParams() url.Values {
	p := o.Order.getParams()

	p.Add("orderId", o.orderId)

	if o.qty != nil {
		p.Add("qty", *o.qty)
	}

	return p
}
