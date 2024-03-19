package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/position"
	"github.com/yasseldg/bybit/rest/response"
)

type GetPositionInfo struct {
	c *common.Client
	r *common.Request

	*position.Position
}

type GetPositionInfoResponse struct {
	common.InfoResponse
	Result response.PositionListInfo `json:"result"`
}

func NewGetPositionInfo(c *common.Client, category string) (*GetPositionInfo, error) {
	positions, err := position.NewPosition(category)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/position/list").Get().Signed()

	return &GetPositionInfo{
		c:        c,
		r:        request,
		Position: positions,
	}, nil
}

func (o *GetPositionInfo) Do(ctx context.Context, opts ...common.RequestOption) (*GetPositionInfoResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(GetPositionInfoResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
