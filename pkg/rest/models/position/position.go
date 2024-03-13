package position

import (
	"net/url"

	"github.com/yasseldg/bybit/pkg/rest/errors"

	"github.com/yasseldg/go-simple/types/sInts"
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
func (p *Position) GetParams() url.Values {
	ps := url.Values{}

	ps.Add("category", p.category)

	if p.symbol != nil {
		ps.Add("symbol", *p.symbol)
	}
	if p.baseCoin != nil {
		ps.Add("baseCoin", *p.baseCoin)
	}
	if p.settleCoin != nil {
		ps.Add("settleCoin", *p.settleCoin)
	}
	if p.limit != nil {
		ps.Add("limit", sInts.ToString(int64(*p.limit)))
	}
	if p.cursor != nil {
		ps.Add("cursor", *p.cursor)
	}

	return ps
}
