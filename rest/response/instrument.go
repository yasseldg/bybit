package response

type InstrumentsInfoResult struct {
	Category       string           `json:"category"`
	List           InstrumentsInfos `json:"list"`
	NextPageCursor string           `json:"nextPageCursor"`
}

type InstrumentsInfo struct {
	Symbol             string         `json:"symbol"`
	ContractType       string         `json:"contractType"`
	Status             string         `json:"status"`
	BaseCoin           string         `json:"baseCoin"`
	QuoteCoin          string         `json:"quoteCoin"`
	LaunchTime         string         `json:"launchTime"`
	DeliveryTime       string         `json:"deliveryTime"`
	DeliveryFeeRate    string         `json:"deliveryFeeRate"`
	PriceScale         string         `json:"priceScale"`
	LeverageFilter     LeverageFilter `json:"leverageFilter"`
	PriceFilter        PriceFilter    `json:"priceFilter"`
	LotSizeFilter      LotSizeFilter  `json:"lotSizeFilter"`
	UnifiedMarginTrade bool           `json:"unifiedMarginTrade"`
	FundingInterval    int            `json:"fundingInterval"`
	SettleCoin         string         `json:"settleCoin"`
}
type InstrumentsInfos []InstrumentsInfo

type LeverageFilter struct {
	MinLeverage  string `json:"minLeverage"`
	MaxLeverage  string `json:"maxLeverage"`
	LeverageStep string `json:"leverageStep"`
}

type PriceFilter struct {
	MinPrice string `json:"minPrice"`
	MaxPrice string `json:"maxPrice"`
	TickSize string `json:"tickSize"`
}

type LotSizeFilter struct {
	MaxOrderQty         string `json:"maxOrderQty"`
	MinOrderQty         string `json:"minOrderQty"`
	QtyStep             string `json:"qtyStep"`
	PostOnlyMaxOrderQty string `json:"postOnlyMaxOrderQty"`
}

// "result": {
//     "category": "linear",
//     "list": [
//         {
//             "symbol": "BTCUSDT",
//             "contractType": "LinearPerpetual",
//             "status": "Trading",
//             "baseCoin": "BTC",
//             "quoteCoin": "USDT",
//             "launchTime": "1585526400000",
//             "deliveryTime": "0",
//             "deliveryFeeRate": "",
//             "priceScale": "2",
//             "leverageFilter": {
//                 "minLeverage": "1",
//                 "maxLeverage": "100.00",
//                 "leverageStep": "0.01"
//             },
//             "priceFilter": {
//                 "minPrice": "0.50",
//                 "maxPrice": "999999.00",
//                 "tickSize": "0.50"
//             },
//             "lotSizeFilter": {
//                 "maxOrderQty": "100.000",
//                 "minOrderQty": "0.001",
//                 "qtyStep": "0.001",
//                 "postOnlyMaxOrderQty": "1000.000"
//             },
//             "unifiedMarginTrade": true,
//             "fundingInterval": 480,
//             "settleCoin": "USDT"
//         }
//     ],
//     "nextPageCursor": ""
// },
