package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/trade"
	"github.com/yasseldg/bybit/rest/response"
)

type CancelOrder struct {
	c common.InterClient
	r *common.Request

	*trade.CancelOrder
}

type CancelOrderResponse struct {
	common.InfoResponse
	Result response.OrderResult `json:"result"`
}

func NewCancelOrder(c common.InterClient, category, symbol, orderId string) (*CancelOrder, error) {
	orders, err := trade.NewCancelOrder(category, symbol, orderId)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/order/cancel").Post().Signed()

	return &CancelOrder{
		c:           c,
		r:           request,
		CancelOrder: orders,
	}, nil
}

func (o *CancelOrder) Do(ctx context.Context, opts ...common.RequestOption) (*CancelOrderResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(CancelOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
