package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/account"
)

type SetMarginMode struct {
	c common.InterClient
	r *common.Request

	*account.MarginMode
}

type SetMarginModeResponse struct {
	common.RetResponse
}

func NewSetMarginMode(c common.InterClient) (*SetMarginMode, error) {
	marginMode, err := account.NewMarginMode()
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/account/set-margin-mode").Post().Signed()

	return &SetMarginMode{
		c:          c,
		r:          request,
		MarginMode: marginMode,
	}, nil
}

func (o *SetMarginMode) Do(ctx context.Context, opts ...common.RequestOption) (*SetMarginModeResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(SetMarginModeResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
