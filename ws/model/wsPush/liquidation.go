package wsPush

type LiquidationResp struct {
	Base `json:",inline"`
	Data Liquidation `json:"data"`
}

type Liquidation struct {
	UpdatedTime int64  `json:"updatedTime"`
	Symbol      string `json:"symbol"`
	Side        string `json:"side"`
	Size        string `json:"size"`
	Price       string `json:"price"`
}

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
