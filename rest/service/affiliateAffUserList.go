package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/affiliate"
	"github.com/yasseldg/bybit/rest/response"
)

type GetAffiliateUserList struct {
	c common.InterClient
	r *common.Request

	*affiliate.UserList
}

type GetAffiliateUserListResponse struct {
	common.RetResponse
	Result response.AffiliateUserList `json:"result"`
}

func NewGetAffiliateUserList(c common.InterClient) (*GetAffiliateUserList, error) {
	affiliateUserList, err := affiliate.NewAffiliateUserList()
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/affiliate/aff-user-list").Get().Signed()

	return &GetAffiliateUserList{
		c:        c,
		r:        request,
		UserList: affiliateUserList,
	}, nil
}

func (o *GetAffiliateUserList) Do(ctx context.Context, opts ...common.RequestOption) (*GetAffiliateUserListResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(GetAffiliateUserListResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
