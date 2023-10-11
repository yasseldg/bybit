package config

import (
	"os"

	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/sEnv"
)

type Url string

var Url_Wss, Url_Rest Url

func init() {
	Url_Wss = Url_Base_Wss
	if sConv.GetBool(sEnv.Get("BYBIT_TESTNET", "false")) {
		Url_Wss = Url_Base_Wss_Testnet
	}

	Url_Rest = Url_Base_Rest
	if sConv.GetBool(sEnv.Get("BYBIT_TESTNET", "false")) {
		Url_Rest = Url_Base_Rest_Testnet
	}
}

const (
	// Mainnet
	Url_Base_Wss  = Url("wss://stream.bybit.com/v5")
	Url_Base_Rest = Url("https://api.bybit.com/v5")

	// Testnet
	Url_Base_Wss_Testnet  = Url("wss://stream-testnet.bybit.com/v5")
	Url_Base_Rest_Testnet = Url("https://api-testnet.bybit.com/v5")
)

// credentials
type ApiCreds struct {
	ApiKey     string
	SecretKey  string
	PASSPHRASE string
}

func GetDefaultCredentials() *ApiCreds {
	return &ApiCreds{
		ApiKey:     os.Getenv("BYBIT_API_KEY"),
		SecretKey:  os.Getenv("BYBIT_API_SECRET"),
		PASSPHRASE: os.Getenv("BYBIT_API_PASS"),
	}
}
