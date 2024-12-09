package account

import (
	"github.com/yasseldg/bybit/common"
)

type HedgingMode struct {
	setHedgingMode string
}

func NewHedgingMode(hedging_mode bool) (*HedgingMode, error) {

	setHedgingMode := "OFF"
	if hedging_mode {
		setHedgingMode = "ON"
	}

	return &HedgingMode{
		setHedgingMode: setHedgingMode,
	}, nil
}

// Params returns the order parameters
func (o *HedgingMode) GetParams() common.Params {
	p := common.NewParams()

	p.Set("setHedgingMode", o.setHedgingMode)

	return p
}
