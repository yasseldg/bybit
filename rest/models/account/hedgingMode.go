package account

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"
)

type HedgingMode struct {
	setHedgingMode string
}

func NewHedgingMode(hedging_mode string) (*HedgingMode, error) {
	if hedging_mode != "ON" && hedging_mode != "OFF" {
		return nil, errors.ErrHedgingMode
	}

	return &HedgingMode{
		setHedgingMode: hedging_mode,
	}, nil
}

// Params returns the order parameters
func (o *HedgingMode) GetParams() common.Params {
	p := common.NewParams()

	p.Set("setHedgingMode", o.setHedgingMode)

	return p
}
