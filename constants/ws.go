package constants

type Channel string

const (
	// WebSocket public channel
	Channel_PublicSpot    = Channel("public/spot")    // Spot
	Channel_PublicLinear  = Channel("public/linear")  // USDT and USDC perpetual
	Channel_PublicInverse = Channel("public/inverse") // Inverse contract
	Channel_PublicOption  = Channel("public/option")  // USDC Option

	// WebSocket private channel
	Channel_Private = Channel("private")
)

type Topic string

const (
	Topic_Kline       = Topic("kline")
	Topic_PublicTrade = Topic("publicTrade")
	Topic_Orderbook   = Topic("orderbook")
	Topic_Tickers     = Topic("tickers")
	Topic_Liquidation = Topic("liquidation")
)

const (
	// WebSocket
	WsAuthMethod        = "GET"
	WsAuthPath          = "/user/verify"
	WsOpAuth            = "auth"
	WsOpUnsubscribe     = "unsubscribe"
	WsOpSubscribe       = "subscribe"
	WsOpPing            = "ping"
	WsOpPong            = "pong"
	PingIntervalSecond  = 10
	TimerIntervalSecond = 2
)
