package wsPush

type TradeResp struct {
	Base `json:",inline"`
	Data Trades `json:"data"`
}

type Trade struct {
	Timestamp  int64  `json:"T"`
	Symbol     string `json:"s"`
	Side       string `json:"S"`
	Size       string `json:"v"`
	Price      string `json:"p"`
	Direction  string `json:"L"`
	ID         string `json:"i"`
	BlockTrade bool   `json:"BT"`
}
type Trades []Trade

// Parameter	Type	Comments
// id	string	Message id. Unique field for option
// topic	string	Topic name
// type	string	Data type. snapshot
// ts	number	The timestamp (ms) that the system generates the data
// data	array	Object
// > T	number	The timestamp (ms) that the order is filled
// > s	string	Symbol name
// > S	string	Side of taker. Buy,Sell
// > v	string	Trade size
// > p	string	Trade price
// > L	string	Direction of price change. Unique field for future
// > i	string	Trade ID
// > BT	boolean	Whether it is a block trade order or not
