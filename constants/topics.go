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

func GetTopicPosition(category string) SubscribeTopic {
	switch category {
	case "linear":
		return SubscribeTopic(Topic_Position_Linear)
	case "inverse":
		return SubscribeTopic(Topic_Position_Inverse)
	case "option":
		return SubscribeTopic(Topic_Position_Option)

	default:
		return SubscribeTopic(Topic_Position)
	}
}

func GetTopicOrder(category string) SubscribeTopic {
	switch category {
	case "spot":
		return SubscribeTopic(Topic_Order_Spot)
	case "linear":
		return SubscribeTopic(Topic_Order_Linear)
	case "inverse":
		return SubscribeTopic(Topic_Order_Inverse)
	case "option":
		return SubscribeTopic(Topic_Order_Option)

	default:
		return SubscribeTopic(Topic_Order)
	}
}
