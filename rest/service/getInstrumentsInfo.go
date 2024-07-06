package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/market"
	"github.com/yasseldg/bybit/rest/response"
)

type GetInstrumentsInfo struct {
	c common.InterClient
	r *common.Request

	*market.InstrumentsInfo
}

type GetInstrumentsInfoResponse struct {
	common.RetResponse

	Result response.InstrumentsInfoResult `json:"result"`
}

func NewGetInstrumentsInfo(c common.InterClient, category string) (*GetInstrumentsInfo, error) {
	instrumentsInfo, err := market.NewInstrumentsInfo(category)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/market/instruments-info").Get()

	return &GetInstrumentsInfo{
		c:               c,
		r:               request,
		InstrumentsInfo: instrumentsInfo,
	}, nil
}

func (o *GetInstrumentsInfo) Do(ctx context.Context, opts ...common.RequestOption) (*GetInstrumentsInfoResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(GetInstrumentsInfoResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
