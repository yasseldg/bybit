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

	Topic_Position         = Topic("position")
	Topic_Position_Linear  = Topic("position.linear")
	Topic_Position_Inverse = Topic("position.inverse")
	Topic_Position_Option  = Topic("position.option")

	Topic_Order         = Topic("order")
	Topic_Order_Spot    = Topic("order.spot")
	Topic_Order_Linear  = Topic("order.linear")
	Topic_Order_Inverse = Topic("order.inverse")
	Topic_Order_Option  = Topic("order.option")
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
