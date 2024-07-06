package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/trade"
	"github.com/yasseldg/bybit/rest/response"
)

type AmendOrder struct {
	c common.InterClient
	r *common.Request

	*trade.AmendOrder
}

type AmendOrderResponse struct {
	common.InfoResponse
	Result response.OrderResult `json:"result"`
}

func NewAmendOrder(c common.InterClient, category, symbol, orderId string, prec int) (*AmendOrder, error) {
	order, err := trade.NewAmendOrder(category, symbol, orderId, prec)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/order/amend").Post().Signed()

	return &AmendOrder{
		c:          c,
		r:          request,
		AmendOrder: order,
	}, nil
}

func (o *AmendOrder) Do(ctx context.Context, opts ...common.RequestOption) (*AmendOrderResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(AmendOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
