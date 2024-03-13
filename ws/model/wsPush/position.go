package wsPush

import "fmt"

type PositionResp struct {
	PrivateBase
	Data []Position `json:"data"`
}

type Position struct {
	PositionIdx            int    `json:"positionIdx"`
	TradeMode              int    `json:"tradeMode"`
	RiskId                 int    `json:"riskId"`
	RiskLimitValue         string `json:"riskLimitValue"`
	Symbol                 string `json:"symbol"`
	Side                   string `json:"side"`
	Size                   string `json:"size"`
	EntryPrice             string `json:"entryPrice"`
	Leverage               string `json:"leverage"`
	PositionValue          string `json:"positionValue"`
	PositionBalance        string `json:"positionBalance"`
	MarkPrice              string `json:"markPrice"`
	PositionIM             string `json:"positionIM"`
	PositionMM             string `json:"positionMM"`
	TakeProfit             string `json:"takeProfit"`
	StopLoss               string `json:"stopLoss"`
	TrailingStop           string `json:"trailingStop"`
	UnrealisedPnl          string `json:"unrealisedPnl"`
	CumRealisedPnl         string `json:"cumRealisedPnl"`
	CreatedTime            string `json:"createdTime"`
	UpdatedTime            string `json:"updatedTime"`
	TpslMode               string `json:"tpslMode"`
	LiqPrice               string `json:"liqPrice"`
	BustPrice              string `json:"bustPrice"`
	Category               string `json:"category"`
	PositionStatus         string `json:"positionStatus"`
	AdlRankIndicator       int    `json:"adlRankIndicator"`
	AutoAddMargin          int    `json:"autoAddMargin"`
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`
	Seq                    int64  `json:"seq"`
	IsReduceOnly           bool   `json:"isReduceOnly"`
}

func (p *Position) String() string {
	return fmt.Sprintf("Position: %s .. %s .. size: %s  .. idx: %d  .. value: %s  .. balance: %s  ..  status: %s",
		p.Symbol, p.Side, p.Size, p.PositionIdx, p.PositionValue, p.PositionBalance, p.PositionStatus)
}

// {
//     "id": "1003076014fb7eedb-c7e6-45d6-a8c1-270f0169171a",
//     "topic": "position",
//     "creationTime": 1697682317044,
//     "data": [
//         {
//             "positionIdx": 2,
//             "tradeMode": 0,
//             "riskId": 1,
//             "riskLimitValue": "2000000",
//             "symbol": "BTCUSDT",
//             "side": "",
//             "size": "0",
//             "entryPrice": "0",
//             "leverage": "10",
//             "positionValue": "0",
//             "positionBalance": "0",
//             "markPrice": "28184.5",
//             "positionIM": "0",
//             "positionMM": "0",
//             "takeProfit": "0",
//             "stopLoss": "0",
//             "trailingStop": "0",
//             "unrealisedPnl": "0",
//             "cumRealisedPnl": "-25.06579337",
//             "createdTime": "1694402496913",
//             "updatedTime": "1697682317038",
//             "tpslMode": "Full",
//             "liqPrice": "0",
//             "bustPrice": "",
//             "category": "linear",
//             "positionStatus": "Normal",
//             "adlRankIndicator": 0,
//             "autoAddMargin": 0,
//             "leverageSysUpdatedTime": "",
//             "mmrSysUpdatedTime": "",
//             "seq": 8327597863,
//             "isReduceOnly": false
//         }
//     ]
// }
