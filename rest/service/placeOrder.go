package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/trade"
	"github.com/yasseldg/bybit/rest/response"
)

type PlaceOrder struct {
	c *common.Client
	r *common.Request

	*trade.PlaceOrder
}

type PlaceOrderResponse struct {
	common.InfoResponse
	Result response.OrderResult `json:"result"`
}

func NewPlaceOrder(c *common.Client, category, symbol, side, orderType string, qty float64, prec int) (*PlaceOrder, error) {
	order, err := trade.NewPlaceOrder(category, symbol, side, orderType, qty, prec)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/order/create").Post().Signed()

	return &PlaceOrder{
		c:          c,
		r:          request,
		PlaceOrder: order,
	}, nil
}

func (o *PlaceOrder) Do(ctx context.Context, opts ...common.RequestOption) (*PlaceOrderResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(PlaceOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
