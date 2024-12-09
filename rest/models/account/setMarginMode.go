package account

import "github.com/yasseldg/bybit/common"

type MarginMode struct {
	setMarginMode string
}

func NewMarginMode() (*MarginMode, error) {
	return &MarginMode{
		setMarginMode: MarginMode_CROSS,
	}, nil
}

func (o *MarginMode) Isolated() {
	o.setMarginMode = MarginMode_ISOLATED
}

func (o *MarginMode) Cross() {
	o.setMarginMode = MarginMode_CROSS
}

func (o *MarginMode) Portfolio() {
	o.setMarginMode = MarginMode_PORTFOLIO
}

// Params returns the order parameters
func (o *MarginMode) GetParams() common.Params {
	p := common.NewParams()

	p.Set("setMarginMode", o.setMarginMode)

	return p
}

const (
	MarginMode_ISOLATED  = "ISOLATED_MARGIN"
	MarginMode_CROSS     = "REGULAR_MARGIN"
	MarginMode_PORTFOLIO = "PORTFOLIO_MARGIN"
)
