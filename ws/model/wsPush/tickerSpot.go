package wsPush

import "github.com/yasseldg/bybit/constants"

type TickerSpotResp struct {
	Base     `json:",inline"`
	CrossSeq int64      `json:"cs,omitempty"`
	ID       string     `json:"id,omitempty"`
	Data     TickerSpot `json:"data"`
}

type TickerSpot struct {
	Symbol        constants.Symbol `json:"symbol"`
	LastPrice     string           `json:"lastPrice"`
	HighPrice24H  string           `json:"highPrice24h"`
	LowPrice24H   string           `json:"lowPrice24h"`
	PrevPrice24H  string           `json:"prevPrice24h"`
	Volume24H     string           `json:"volume24h"`
	Turnover24H   string           `json:"turnover24h"`
	Price24HPcnt  string           `json:"price24hPcnt"`
	UsdIndexPrice string           `json:"usdIndexPrice"`
}

// Parameter	Type	Comments
// topic	string	Topic name
// ts	number	The timestamp (ms) that the system generates the data
// type	string	Data type. snapshot
// cs	integer	Cross sequence
// data	array	Object
// > symbol	string	Symbol name
// > lastPrice	string	Last price
// > highPrice24h	string	The highest price in the last 24 hours
// > lowPrice24h	string	The lowest price in the last 24 hours
// > prevPrice24h	string	Percentage change of market price relative to 24h
// > volume24h	string	Volume for 24h
// > turnover24h	string	Turnover for 24h
// > price24hPcnt	string	Percentage change of market price relative to 24h
// > usdIndexPrice	string	USD index price. It can be empty
