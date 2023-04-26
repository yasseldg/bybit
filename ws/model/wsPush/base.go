package wsPush

type Base struct {
	Topic     string `json:"topic"`
	Type      string `json:"type"`
	TimeStamp int64  `json:"ts"`
	// Data   []interface{} `json:"data"`
}
