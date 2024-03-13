package bybit

import (
	"github.com/yasseldg/bybit/constants"
	"github.com/yasseldg/bybit/ws/common"

	"github.com/yasseldg/go-simple/logs/sLog"
)

// ws
type BybitWsClient struct {
	WsClient *common.WsClient
}

func NewWsClient(channel constants.Channel, needLogin bool) *BybitWsClient {
	wsc := new(common.WsClient)
	wsc.Init(channel, needLogin, func(message string) {
		sLog.Debug("WebSocket message:" + message)
	}, func(message string) {
		sLog.Error(message)
	})
	return &BybitWsClient{wsc}
}

func WscPublicSpot() *BybitWsClient {
	return NewWsClient(constants.Channel_PublicSpot, false)
}

func WscPublicLinear() *BybitWsClient {
	return NewWsClient(constants.Channel_PublicLinear, false)
}

func WscPublicInverse() *BybitWsClient {
	return NewWsClient(constants.Channel_PublicInverse, false)
}

func WscPublicOption() *BybitWsClient {
	return NewWsClient(constants.Channel_PublicOption, false)
}

func WscPrivate() *BybitWsClient {
	return NewWsClient(constants.Channel_Private, true)
}

type UnsuscribeFunc func()

func (bws *BybitWsClient) Subscribe(listener common.OnReceive, topics ...constants.SubscribeTopic) UnsuscribeFunc {
	bws.WsClient.Subscribe(topics, listener)
	return func() { bws.WsClient.UnSubscribe(topics) }
}

func (bws *BybitWsClient) SubscribeOrderbook(listener common.OnReceive, symbol constants.Symbol, interval constants.Interval) UnsuscribeFunc {
	return bws.Subscribe(listener, constants.GetTopicOrderbook(symbol, interval))
}

func (bws *BybitWsClient) SubscribeKline(listener common.OnReceive, symbol constants.Symbol, interval constants.Interval) UnsuscribeFunc {
	return bws.Subscribe(listener, constants.GetTopicKline(symbol, interval))
}

func (bws *BybitWsClient) SubscribeTrade(listener common.OnReceive, symbol constants.Symbol) UnsuscribeFunc {
	return bws.Subscribe(listener, constants.GetTopicTrade(symbol))
}

func (bws *BybitWsClient) SubscribeTickers(listener common.OnReceive, symbol constants.Symbol) UnsuscribeFunc {
	return bws.Subscribe(listener, constants.GetTopicTickers(symbol))
}

func (bws *BybitWsClient) SubscribeLiquidation(listener common.OnReceive, symbol constants.Symbol) UnsuscribeFunc {
	return bws.Subscribe(listener, constants.GetTopicLiquidation(symbol))
}
