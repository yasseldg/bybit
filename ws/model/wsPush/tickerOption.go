package wsPush

import "github.com/yasseldg/bybit/constants"

type TickerOptionResp struct {
	Base     `json:",inline"`
	CrossSeq int64        `json:"cs,omitempty"`
	ID       string       `json:"id,omitempty"`
	Data     TickerOption `json:"data"`
}

type TickerOption struct {
	Symbol                 constants.Symbol `json:"symbol"`
	BidPrice               string           `json:"bidPrice"`
	BidSize                string           `json:"bidSize"`
	BidIv                  string           `json:"bidIv"`
	AskPrice               string           `json:"askPrice"`
	AskSize                string           `json:"askSize"`
	AskIv                  string           `json:"askIv"`
	LastPrice              string           `json:"lastPrice"`
	HighPrice24H           string           `json:"highPrice24h"`
	LowPrice24H            string           `json:"lowPrice24h"`
	MarkPrice              string           `json:"markPrice"`
	IndexPrice             string           `json:"indexPrice"`
	MarkPriceIv            string           `json:"markPriceIv"`
	UnderlyingPrice        string           `json:"underlyingPrice"`
	OpenInterest           string           `json:"openInterest"`
	Turnover24H            string           `json:"turnover24h"`
	Volume24H              string           `json:"volume24h"`
	TotalVolume            string           `json:"totalVolume"`
	TotalTurnover          string           `json:"totalTurnover"`
	Delta                  string           `json:"delta"`
	Gamma                  string           `json:"gamma"`
	Vega                   string           `json:"vega"`
	Theta                  string           `json:"theta"`
	PredictedDeliveryPrice string           `json:"predictedDeliveryPrice"`
	Change24H              string           `json:"change24h"`
}

// Parameter	Type	Comments
// topic	string	Topic name
// type	string	Data type. snapshot
// id	string	message ID
// ts	number	The timestamp (ms) that the system generates the data
// data	array	Object
// > symbol	string	Symbol name
// > bidPrice	string	Best bid price
// > bidSize	string	Best bid size
// > bidIv	string	Best bid iv
// > askPrice	string	Best ask price
// > askSize	string	Best ask size
// > askIv	string	Best ask iv
// > lastPrice	string	Last price
// > highPrice24h	string	The highest price in the last 24 hours
// > lowPrice24h	string	The lowest price in the last 24 hours
// > markPrice	string	Mark price
// > indexPrice	string	Index price
// > markPriceIv	string	Mark price iv
// > underlyingPrice	string	Underlying price
// > openInterest	string	Open interest size
// > turnover24h	string	Turnover for 24h
// > volume24h	string	Volume for 24h
// > totalVolume	string	Total volume
// > totalTurnover	string	Total turnover
// > delta	string	Delta
// > gamma	string	Gamma
// > vega	string	Vega
// > theta	string	Theta
// > predictedDeliveryPrice	string	Predicated delivery price. It has value when 30 min before delivery
// > change24h	string	The change in the last 24 hous
