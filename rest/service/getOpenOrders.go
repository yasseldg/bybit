package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/trade"
	"github.com/yasseldg/bybit/rest/response"
)

type GetOpenOrders struct {
	c *common.Client
	r *common.Request

	*trade.GetOpenOrders
}

type GetOpenOrdersResponse struct {
	common.InfoResponse
	Result response.OpenOrdersInfo `json:"result"`
}

func NewGetOpenOrders(c *common.Client, category string) (*GetOpenOrders, error) {
	orders, err := trade.NewGetOpenOrders(category)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/order/realtime").Get().Signed()

	return &GetOpenOrders{
		c:             c,
		r:             request,
		GetOpenOrders: orders,
	}, nil
}

func (o *GetOpenOrders) Do(ctx context.Context, opts ...common.RequestOption) (*GetOpenOrdersResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(GetOpenOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
