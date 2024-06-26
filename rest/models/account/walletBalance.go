package account

import (
	"github.com/yasseldg/bybit/common"
)

type WalletBalance struct {
	accountType string
	coins       []string
}

func NewWalletBalance(accountType string) (*WalletBalance, error) {
	return &WalletBalance{
		accountType: accountType,
	}, nil
}

func (wb *WalletBalance) Coins(coins []string) *WalletBalance {
	wb.coins = append(wb.coins, coins...)
	return wb
}

// Params returns the order parameters
func (wb *WalletBalance) GetParams() common.Params {
	p := common.NewParams()

	p.Set("accountType", wb.accountType)

	if len(wb.coins) > 0 {
		coins := ""
		for i, coin := range wb.coins {
			if i > 0 {
				coins += ","
			}
			coins += coin
		}

		p.Set("coin", coins)
	}

	return p
}
