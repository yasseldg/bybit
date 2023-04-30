package wsPush

import "github.com/yasseldg/bybit/constants"

type TickerResp struct {
	Base     `json:",inline"`
	CrossSeq int64  `json:"cs,omitempty"`
	ID       string `json:"id,omitempty"`
	Data     Ticker `json:"data"`
}

type Ticker struct {
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
	OpenInterestValue      string           `json:"openInterestValue"`
	Turnover24H            string           `json:"turnover24h"`
	Volume24H              string           `json:"volume24h"`
	FundingRate            string           `json:"fundingRate"`
	NextFundingTime        string           `json:"nextFundingTime"`
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
// type	string	Data type. snapshot,delta
// cs	integer	Cross sequence
// ts	number	The timestamp (ms) that the system generates the data
// data	array	Object
// > symbol	string	Symbol name
// > tickDirection	string	Tick direction
// > price24hPcnt	string	Percentage change of market price in the last 24 hours
// > lastPrice	string	Last price
// > prevPrice24h	string	Market price 24 hours ago
// > highPrice24h	string	The highest price in the last 24 hours
// > lowPrice24h	string	The lowest price in the last 24 hours
// > prevPrice1h	string	Market price an hour ago
// > markPrice	string	Mark price
// > indexPrice	string	Index price
// > openInterest	string	Open interest size
// > openInterestValue	string	Open interest value
// > turnover24h	string	Turnover for 24h
// > volume24h	string	Volume for 24h
// > nextFundingTime	string	Next funding timestamp (ms)
// > fundingRate	string	Funding rate
// > bid1Price	string	Best bid price
// > bid1Size	string	Best bid size
// > ask1Price	string	Best ask price
// > ask1Size	string	Best ask size
// > deliveryTime	datetime	Delivery date time (UTC+0). Unique field for inverse futures
// > basisRate	string	Basis rate. Unique field for inverse futures
// > deliveryFeeRate	string	Delivery fee rate. Unique field for inverse futures
// > predictedDeliveryPrice	string	Predicated delivery price. Unique field for inverse futures

// Topic:   ..  Type:   ..  TimeStamp: 0  ..  data: {"topic":"tickers.ETHUSDT","type":"snapshot","data":{"symbol":"ETHUSDT","tickDirection":"MinusTick","price24hPcnt":"0.074878","lastPrice":"1951.28","prevPrice24h":"1815.35","highPrice24h":"1964.56","lowPrice24h":"1808.53","prevPrice1h":"1942.50","markPrice":"1951.53","indexPrice":"1952.37","openInterest":"451956.29","openInterestValue":"882006258.62","turnover24h":"2222368225.2433","volume24h":"1176494.5100","nextFundingTime":"1682524800000","fundingRate":"0.0001","bid1Price":"1951.28","bid1Size":"242.76","ask1Price":"1951.29","ask1Size":"32.23"},"cs":65473881041,"ts":1682524024575}
