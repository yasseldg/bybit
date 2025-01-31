package config

import (
	"github.com/yasseldg/go-simple/configs/sEnv"
)

type Url string

var (
	Url_Wss, Url_Rest Url
)

func init() {
	Url_Wss = Url_Base_Wss
	if sEnv.GetBool("BYBIT_TESTNET", false) {
		Url_Wss = Url_Base_Wss_Testnet
	}

	Url_Rest = Url_Base_Rest
	if sEnv.GetBool("BYBIT_TESTNET", false) {
		Url_Rest = Url_Base_Rest_Testnet
	}
}

const (
	// Mainnet
	Url_Base_Wss  = Url("wss://stream.bybit.com/v5")
	Url_Base_Rest = Url("https://api.bybit.com")

	// Testnet
	Url_Base_Wss_Testnet  = Url("wss://stream-testnet.bybit.com/v5")
	Url_Base_Rest_Testnet = Url("https://api-testnet.bybit.com")
)
