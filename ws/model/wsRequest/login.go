package wsRequest

type LoginBase struct {
	Op   string `json:"op"`
	Args Logins `json:"args"`
}

type Login struct {
	ApiKey     string `json:"apiKey"`
	Passphrase string `json:"passphrase"`
	Timestamp  string `json:"timestamp"`
	Sign       string `json:"sign"`
}
type Logins []Login
