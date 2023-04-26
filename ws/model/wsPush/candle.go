package wsPush

import "github.com/yasseldg/bybit/constants"

type CandleResp struct {
	Base `json:",inline"`
	Data Candles `json:"data"`
}

type Candle struct {
	Start     int64              `json:"start"`
	End       int64              `json:"end"`
	Interval  constants.Interval `json:"interval"`
	Open      string             `json:"open"`
	Close     string             `json:"close"`
	High      string             `json:"high"`
	Low       string             `json:"low"`
	Volume    string             `json:"volume"`
	Turnover  string             `json:"turnover"`
	Confirm   bool               `json:"confirm"`
	Timestamp int64              `json:"timestamp"`
}
type Candles []Candle

// Parameter	Type	Comments
// topic	string	Topic name
// type	string	Data type. snapshot
// ts	number	The timestamp (ms) that the system generates the data
// data	array	Object
// > start	number	The start timestamp (ms)
// > end	number	The end timestamp (ms). It is current timestamp if it does not reach to the end time of candle
// > interval	string	Kline interval
// > open	string	Open price
// > close	string	Close price
// > high	string	Highest price
// > low	string	Lowest price
// > volume	string	Trade volume
// > turnover	string	Turnover
// > confirm	boolean	Weather the tick is ended or not
// > timestamp	number	The timestamp (ms) of the last matched order in the candle
