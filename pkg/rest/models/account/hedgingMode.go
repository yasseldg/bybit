package account

import (
	"net/url"

	"github.com/yasseldg/bybit/pkg/rest/errors"
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
func (o *HedgingMode) GetParams() url.Values {
	p := url.Values{}

	p.Add("setHedgingMode", o.setHedgingMode)

	return p
}
