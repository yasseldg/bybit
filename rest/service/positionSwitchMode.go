package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/position"
)

type SwitchPositionMode struct {
	c common.InterClient
	r *common.Request

	*position.SwitchPositionMode
}

type SwitchPositionModeResponse struct {
	common.RetResponse
}

func NewSwitchPositionMode(c common.InterClient, category string) (*SwitchPositionMode, error) {
	switchPositionMode, err := position.NewSwitchPositionMode(category)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/position/switch-mode").Post().Signed()

	return &SwitchPositionMode{
		c:                  c,
		r:                  request,
		SwitchPositionMode: switchPositionMode,
	}, nil
}

func (o *SwitchPositionMode) Do(ctx context.Context, opts ...common.RequestOption) (*SwitchPositionModeResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(SwitchPositionModeResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
