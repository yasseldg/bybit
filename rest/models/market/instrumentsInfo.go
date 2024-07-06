package market

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"
)

type InstrumentsInfo struct {
	category string
	symbol   *string
	status   *string
	baseCoin *string
	limit    *int
	cursor   *string
}

func NewInstrumentsInfo(category string) (*InstrumentsInfo, error) {
	if len(category) == 0 {
		return nil, errors.ErrInvalidCategory
	}

	return &InstrumentsInfo{
		category: category,
	}, nil
}

func (p *InstrumentsInfo) Symbol(symbol string) *InstrumentsInfo {
	p.symbol = &symbol
	return p
}

func (p *InstrumentsInfo) BaseCoin(baseCoin string) *InstrumentsInfo {
	p.baseCoin = &baseCoin
	return p
}

func (p *InstrumentsInfo) Status(status string) *InstrumentsInfo {
	p.status = &status
	return p
}

func (p *InstrumentsInfo) Limit(limit int) *InstrumentsInfo {
	p.limit = &limit
	return p
}

func (p *InstrumentsInfo) Cursor(cursor string) *InstrumentsInfo {
	p.cursor = &cursor
	return p
}

// Params returns the order parameters
func (p *InstrumentsInfo) GetParams() common.Params {
	cp := common.NewParams()

	cp.Set("category", p.category)

	if p.symbol != nil {
		cp.Set("symbol", *p.symbol)
	}
	if p.baseCoin != nil {
		cp.Set("baseCoin", *p.baseCoin)
	}
	if p.status != nil {
		cp.Set("status", *p.status)
	}
	if p.limit != nil {
		cp.Set("limit", *p.limit)
	}
	if p.cursor != nil {
		cp.Set("cursor", *p.cursor)
	}

	return cp
}

// Parameter	Required	Type	Comments
// category	true	string	Product type. spot,linear,inverse,option
// symbol	false	string	Symbol name, like BTCUSDT, uppercase only
// status	false	string	Symbol status filter

//     spot has Trading only
//     PreLaunch: when category=linear&status=PreLaunch, it returns pre-market perpetual contract

// baseCoin	false	string	Base coin, uppercase only
// Apply tolinear,inverse,option only
// option: it returns BTC by default
// limit	false	integer	Limit for data size per page. [1, 1000]. Default: 500
// cursor	false	string	Cursor. Use the nextPageCursor token from the response to retrieve the next page of the result set
