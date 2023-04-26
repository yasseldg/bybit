package wsRequest

type Base struct {
	Op   string        `json:"op"`
	Args []interface{} `json:"args"`
}
