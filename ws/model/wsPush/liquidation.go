package wsPush

type LiquidationResp struct {
	Base `json:",inline"`
	Data Liquidations `json:"data"`
}

type Liquidation struct {
	UpdateTime int64  `json:"updateTime"`
	Symbol     string `json:"symbol"`
	Side       string `json:"side"`
	Size       string `json:"size"`
	Price      string `json:"price"`
}
type Liquidations []Liquidation

// Parameter	Type	Comments
// topic	string	Topic name
// type	string	Data type. snapshot
// ts	number	The timestamp (ms) that the system generates the data
// data	array	Object
// > updateTime	number	The updated timestamp (ms)
// > symbol	string	Symbol name
// > side	string	Order side. Buy,Sell
// > size	string	Executed size
// > price	string	Executed price
