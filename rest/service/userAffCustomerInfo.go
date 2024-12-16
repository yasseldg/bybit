package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/user"
	"github.com/yasseldg/bybit/rest/response"
)

type GetAffiliateUserInfo struct {
	c common.InterClient
	r *common.Request

	*user.AffiliateUserInfo
}

type GetAffiliateUserInfoResponse struct {
	common.RetResponse
	Result response.AffiliateUserInfo `json:"result"`
}

func NewGetAffiliateUserInfo(c common.InterClient, uid string) (*GetAffiliateUserInfo, error) {
	affiliateUserInfo, err := user.NewAffiliateUserInfo(uid)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/user/aff-customer-info").Get().Signed()

	return &GetAffiliateUserInfo{
		c:                 c,
		r:                 request,
		AffiliateUserInfo: affiliateUserInfo,
	}, nil
}

func (o *GetAffiliateUserInfo) Do(ctx context.Context, opts ...common.RequestOption) (*GetAffiliateUserInfoResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(GetAffiliateUserInfoResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
