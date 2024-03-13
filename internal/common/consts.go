package common

const (
	Name    = "bybit.api.go"
	Version = "1.0.1"

	// Rest
	REST_MAINNET       = "https://api.bybit.com"
	REST_MAINNET_BACKT = "https://api.bytick.com"
	REST_TESTNET       = "https://api-testnet.bybit.com"

	// Globals
	timestampKey  = "X-BAPI-TIMESTAMP"
	signatureKey  = "X-BAPI-SIGN"
	apiRequestKey = "X-BAPI-API-KEY"
	recvWindowKey = "X-BAPI-RECV-WINDOW"
	signTypeKey   = "X-BAPI-SIGN-TYPE"

	recvWindow = "5000"

	// WebSocket
	WS_MAINNET = "wss://stream.bybit.com"
	WS_TESTNET = "wss://stream-testnet.bybit.com"
)
