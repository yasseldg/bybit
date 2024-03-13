package constants

import (
	"fmt"
)

type SubscribeTopic string

// orderbook.1.BTCUSDT
func GetTopicOrderbook(symbol Symbol, interval Interval) SubscribeTopic {
	return SubscribeTopic(fmt.Sprintf("%s.%s.%s", Topic_Orderbook, interval, symbol))
}

// "kline.5.BTCUSDT"
func GetTopicKline(symbol Symbol, interval Interval) SubscribeTopic {
	return SubscribeTopic(fmt.Sprintf("%s.%s.%s", Topic_Kline, interval, symbol))
}

// publicTrade.BTCUSDT
func GetTopicTrade(symbol Symbol) SubscribeTopic {
	return SubscribeTopic(fmt.Sprintf("%s.%s", Topic_PublicTrade, symbol))
}

// tickers.{symbol}
func GetTopicTickers(symbol Symbol) SubscribeTopic {
	return SubscribeTopic(fmt.Sprintf("%s.%s", Topic_Tickers, symbol))
}

// liquidation.BTCUSDT
func GetTopicLiquidation(symbol Symbol) SubscribeTopic {
	return SubscribeTopic(fmt.Sprintf("%s.%s", Topic_Liquidation, symbol))
}

func GetTopicPositionLinear() SubscribeTopic {
	return SubscribeTopic(Topic_Position_Linear)
}

func GetTopicOrderLinear() SubscribeTopic {
	return SubscribeTopic(Topic_Order_Linear)
}
