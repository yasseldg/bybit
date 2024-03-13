package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/pkg/rest/models/account"

	"github.com/yasseldg/bybit/internal/common"
)

type SetSpotHedging struct {
	c *common.Client
	r *common.Request

	*account.HedgingMode
}

type SetSpotHedgingResponse struct {
	common.RetResponse
}

func NewSetSpotHedging(c *common.Client, hedging_mode string) (*SetSpotHedging, error) {
	hedgingMode, err := account.NewHedgingMode(hedging_mode)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/account/set-hedging-mode").Post().Signed()

	return &SetSpotHedging{
		c:           c,
		r:           request,
		HedgingMode: hedgingMode,
	}, nil
}

func (o *SetSpotHedging) Do(ctx context.Context, opts ...common.RequestOption) (*SetSpotHedgingResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(SetSpotHedgingResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
