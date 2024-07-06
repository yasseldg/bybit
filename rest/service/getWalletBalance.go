package service

import (
	"context"
	"encoding/json"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/models/account"
	"github.com/yasseldg/bybit/rest/response"
)

type GetWalletBalance struct {
	c common.InterClient
	r *common.Request

	*account.WalletBalance
}

type GetWalletBalanceResponse struct {
	common.RetResponse
	Result response.WalletBalanceInfo `json:"result"`
}

func NewGetWalletBalance(c common.InterClient, accountType string) (*GetWalletBalance, error) {
	walletBalance, err := account.NewWalletBalance(accountType)
	if err != nil {
		return nil, err
	}

	request := new(common.Request).EndPoint("/v5/account/wallet-balance").Get().Signed()

	return &GetWalletBalance{
		c: c,
		r: request,

		WalletBalance: walletBalance,
	}, nil
}

func (o *GetWalletBalance) Do(ctx context.Context, opts ...common.RequestOption) (*GetWalletBalanceResponse, error) {

	o.r.SetParams(o.GetParams())

	data, err := o.c.Call(ctx, o.r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(GetWalletBalanceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
