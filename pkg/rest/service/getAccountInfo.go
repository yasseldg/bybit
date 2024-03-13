package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/pkg/rest/response"

	"github.com/yasseldg/bybit/internal/common"
)

type GetAccountInfo struct {
	c *common.Client
	r *common.Request
}

type GetAccountInfoResponse struct {
	common.RetResponse
	Result response.AccountInfo `json:"result"`
}

func NewGetAccountInfo(c *common.Client, hedging_mode string) (*GetAccountInfo, error) {

	request := new(common.Request).EndPoint("/v5/account/info").Get().Signed()

	return &GetAccountInfo{
		c: c,
		r: request,
	}, nil
}

func (o *GetAccountInfo) Do(ctx context.Context, opts ...common.RequestOption) (*GetAccountInfoResponse, error) {

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(GetAccountInfoResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
