package restResponse

type TickersResp struct {
	Base   `json:",inline"`
	Result TickersResult `json:"result"`
}

type TickersResult struct {
	Category string   `json:"category"`
	List     Tickerss `json:"list"`
}

type Tickers struct {
	Symbol                 string `json:"symbol"`
	LastPrice              string `json:"lastPrice"`
	IndexPrice             string `json:"indexPrice"`
	MarkPrice              string `json:"markPrice"`
	PrevPrice24h           string `json:"prevPrice24h"`
	Price24hPcnt           string `json:"price24hPcnt"`
	HighPrice24h           string `json:"highPrice24h"`
	LowPrice24h            string `json:"lowPrice24h"`
	PrevPrice1h            string `json:"prevPrice1h"`
	OpenInterest           string `json:"openInterest"`
	OpenInterestValue      string `json:"openInterestValue"`
	Turnover24h            string `json:"turnover24h"`
	Volume24h              string `json:"volume24h"`
	FundingRate            string `json:"fundingRate"`
	NextFundingTime        string `json:"nextFundingTime"`
	PredictedDeliveryPrice string `json:"predictedDeliveryPrice"`
	BasisRate              string `json:"basisRate"`
	DeliveryFeeRate        string `json:"deliveryFeeRate"`
	DeliveryTime           string `json:"deliveryTime"`
	Ask1Size               string `json:"ask1Size"`
	Bid1Price              string `json:"bid1Price"`
	Ask1Price              string `json:"ask1Price"`
	Bid1Size               string `json:"bid1Size"`
	Basis                  string `json:"basis"`
}
type Tickerss []Tickers

//     "result": {
//         "category": "inverse",
//         "list": [
//             {
//                 "symbol": "BTCUSD",
//                 "lastPrice": "16597.00",
//                 "indexPrice": "16598.54",
//                 "markPrice": "16596.00",
//                 "prevPrice24h": "16464.50",
//                 "price24hPcnt": "0.008047",
//                 "highPrice24h": "30912.50",
//                 "lowPrice24h": "15700.00",
//                 "prevPrice1h": "16595.50",
//                 "openInterest": "373504107",
//                 "openInterestValue": "22505.67",
//                 "turnover24h": "2352.94950046",
//                 "volume24h": "49337318",
//                 "fundingRate": "-0.001034",
//                 "nextFundingTime": "1672387200000",
//                 "predictedDeliveryPrice": "",
//                 "basisRate": "",
//                 "deliveryFeeRate": "",
//                 "deliveryTime": "0",
//                 "ask1Size": "1",
//                 "bid1Price": "16596.00",
//                 "ask1Price": "16597.50",
//                 "bid1Size": "1",
//                 "basis": ""
//             }
//         ]
//     },
