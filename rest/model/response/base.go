package restResponse

type Base struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int64       `json:"time"`
	// Result   []interface{} `json:"result"`
}
