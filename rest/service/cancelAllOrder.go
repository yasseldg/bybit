package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/trade"
	"github.com/yasseldg/bybit/rest/response"
)

type CancelAllOrders struct {
	c *common.Client
	r *common.Request

	*trade.CancelAllOrders
}

type CancelAllOrdersResponse struct {
	common.InfoResponse
	Result response.ListOrderResult `json:"result"`
}

func NewCancelAllOrders(c *common.Client, category string) (*CancelAllOrders, error) {
	orders, err := trade.NewCancelAllOrders(category)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/order/cancel-all").Post().Signed()

	return &CancelAllOrders{
		c:               c,
		r:               request,
		CancelAllOrders: orders,
	}, nil
}

func (o *CancelAllOrders) Do(ctx context.Context, opts ...common.RequestOption) (*CancelAllOrdersResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(CancelAllOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
