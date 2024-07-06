package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/response"
)

type GetAPIKeyInfo struct {
	c common.InterClient
	r *common.Request
}

type GetAPIKeyInfoResponse struct {
	common.RetResponse
	Result response.UserAPIKeyInfo `json:"result"`
}

func NewGetApiKeyInfo(c common.InterClient) (*GetAPIKeyInfo, error) {

	request := new(common.Request).EndPoint("/v5/user/query-api").Get().Signed()

	return &GetAPIKeyInfo{
		c: c,
		r: request,
	}, nil
}

func (o *GetAPIKeyInfo) Do(ctx context.Context, opts ...common.RequestOption) (*GetAPIKeyInfoResponse, error) {

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(GetAPIKeyInfoResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
