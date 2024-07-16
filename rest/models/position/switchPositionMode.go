package position

import (
	"strings"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"
)

type SwitchPositionMode struct {
	category string
	symbol   *string
	coin     *string
	mode     *int
}

func NewSwitchPositionMode(category string) (*SwitchPositionMode, error) {
	if len(category) == 0 {
		return nil, errors.ErrInvalidCategory
	}

	return &SwitchPositionMode{
		category: category,
	}, nil
}

func (p *SwitchPositionMode) Symbol(symbol string) *SwitchPositionMode {
	symbol = strings.ToUpper(symbol)
	p.symbol = &symbol
	return p
}

func (p *SwitchPositionMode) Coin(coin string) *SwitchPositionMode {
	coin = strings.ToUpper(coin)
	p.coin = &coin
	return p
}

func (p *SwitchPositionMode) Mode(mode int) *SwitchPositionMode {
	p.mode = &mode
	return p
}

// Params returns the order parameters
func (p *SwitchPositionMode) GetParams() common.Params {
	cp := common.NewParams()

	cp.Set("category", p.category)

	if p.symbol != nil {
		cp.Set("symbol", *p.symbol)
	}
	if p.coin != nil {
		cp.Set("coin", *p.coin)
	}
	if p.mode != nil {
		cp.Set("mode", *p.mode)
	}

	return cp
}
