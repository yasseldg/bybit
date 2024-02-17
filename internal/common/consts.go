package common

const (
	Name    = "bybit.api.go"
	Version = "1.0.1"

	// Https
	MAINNET       = "https://api.bybit.com"
	MAINNET_BACKT = "https://api.bytick.com"
	TESTNET       = "https://api-testnet.bybit.com"

	// Globals
	timestampKey  = "X-BAPI-TIMESTAMP"
	signatureKey  = "X-BAPI-SIGN"
	apiRequestKey = "X-BAPI-API-KEY"
	recvWindowKey = "X-BAPI-RECV-WINDOW"
	signTypeKey   = "X-BAPI-SIGN-TYPE"

	recvWindow = "5000"
)
