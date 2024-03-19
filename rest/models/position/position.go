package position

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/errors"
)

type Position struct {
	category   string
	symbol     *string
	baseCoin   *string
	settleCoin *string
	limit      *int
	cursor     *string
}

func NewPosition(category string) (*Position, error) {
	if len(category) == 0 {
		return nil, errors.ErrInvalidCategory
	}

	return &Position{
		category: category,
	}, nil
}

func (p *Position) Symbol(symbol string) *Position {
	p.symbol = &symbol
	return p
}

func (p *Position) BaseCoin(baseCoin string) *Position {
	p.baseCoin = &baseCoin
	return p
}

func (p *Position) SettleCoin(settleCoin string) *Position {
	p.settleCoin = &settleCoin
	return p
}

func (p *Position) Limit(limit int) *Position {
	p.limit = &limit
	return p
}

func (p *Position) Cursor(cursor string) *Position {
	p.cursor = &cursor
	return p
}

// Params returns the order parameters
func (p *Position) GetParams() common.Params {
	cp := common.NewParams()

	cp.Set("category", p.category)

	if p.symbol != nil {
		cp.Set("symbol", *p.symbol)
	}
	if p.baseCoin != nil {
		cp.Set("baseCoin", *p.baseCoin)
	}
	if p.settleCoin != nil {
		cp.Set("settleCoin", *p.settleCoin)
	}
	if p.limit != nil {
		cp.Set("limit", *p.limit)
	}
	if p.cursor != nil {
		cp.Set("cursor", *p.cursor)
	}

	return cp
}
