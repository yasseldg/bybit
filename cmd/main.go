package main

import (
	"fmt"
	"time"

	"github.com/yasseldg/bybit"
	"github.com/yasseldg/bybit/constants"
	"github.com/yasseldg/bybit/ws/model/wsPush"

	"github.com/yasseldg/simplego/sJson"
	"github.com/yasseldg/simplego/sLog"
)

func main() {
	t := time.Now()
	sLog.Info("Cmd Main \n\n")

	// web socket
	wssLinear() //  comment for test Inverse
	wssInverse()

	fmt.Println()
	sLog.Info("Cmd Main: time: %d Segundos \n\n", time.Since(t)/time.Second)
}

func wssLinear() {

	wsc := bybit.WscPublicLinear()

	uFunc_1 := wsc.Subscribe(listCandle, constants.GetTopicKline(constants.Symbol_BTCUSDT, constants.Interval_1m))
	defer uFunc_1()

	uFunc_5 := wsc.SubscribeKline(listCandle, constants.Symbol_ETHUSDT, constants.Interval_5m)
	defer uFunc_5()

	uFunc_t1 := wsc.SubscribeTrade(listTrade, constants.Symbol_ETHUSDT)
	defer uFunc_t1()

	uFunc_t2 := wsc.SubscribeTrade(listTrade, constants.Symbol_BTCUSDT)
	defer uFunc_t2()

	uFunc_l1 := wsc.SubscribeLiquidation(listLiquidation, constants.Symbol_ETHUSDT)
	defer uFunc_l1()

	uFunc_l2 := wsc.SubscribeLiquidation(listLiquidation, constants.Symbol_BTCUSDT)
	defer uFunc_l2()

	uFunc_tr1 := wsc.SubscribeTickers(listTickers, constants.Symbol_ETHUSDT)
	defer uFunc_tr1()

	uFunc_tr2 := wsc.SubscribeTickers(listTickers, constants.Symbol_BTCUSDT)
	defer uFunc_tr2()

	forever := make(chan struct{})
	<-forever
}

func wssInverse() {

	wsc := bybit.WscPublicInverse()

	uFunc_1 := wsc.Subscribe(listCandle, constants.GetTopicKline(constants.Symbol_LTCUSD, constants.Interval_1m))
	defer uFunc_1()

	uFunc_5 := wsc.SubscribeKline(listCandle, constants.Symbol_ETHUSD, constants.Interval_5m)
	defer uFunc_5()

	uFunc_t1 := wsc.SubscribeTrade(listTrade, constants.Symbol_ETHUSD)
	defer uFunc_t1()

	uFunc_t2 := wsc.SubscribeTrade(listTrade, constants.Symbol_BTCUSD)
	defer uFunc_t2()

	uFunc_l1 := wsc.SubscribeLiquidation(listLiquidation, constants.Symbol_ETHUSD)
	defer uFunc_l1()

	uFunc_l2 := wsc.SubscribeLiquidation(listLiquidation, constants.Symbol_BTCUSD)
	defer uFunc_l2()

	uFunc_tr1 := wsc.SubscribeTickers(listTickers, constants.Symbol_ETHUSD)
	defer uFunc_tr1()

	uFunc_tr2 := wsc.SubscribeTickers(listTickers, constants.Symbol_BTCUSD)
	defer uFunc_tr2()

	forever := make(chan struct{})
	<-forever
}

func listCandle(msg string) {
	var pushObj wsPush.CandleResp
	err := sJson.ToObj(msg, &pushObj)
	if err != nil {
		sLog.Error("listCandle: sJson.ToObj(): %s", err)
		return
	}

	// sLog.Debug("Topic: %s  ..  Type: %s  ..  TimeStamp: %d", pushObj.Topic, pushObj.Type, pushObj.TimeStamp)

	for _, data := range pushObj.Data {
		sLog.Debug("%s: %+v ", pushObj.Topic, data)
	}
}

func listTrade(msg string) {
	var pushObj wsPush.TradeResp
	err := sJson.ToObj(msg, &pushObj)
	if err != nil {
		sLog.Error("listTrade: sJson.ToObj(): %s", err)
		return
	}

	// sLog.Debug("Topic: %s  ..  Type: %s  ..  TimeStamp: %d", pushObj.Topic, pushObj.Type, pushObj.TimeStamp)

	for _, data := range pushObj.Data {
		sLog.Debug("%s: %+v ", pushObj.Topic, data)
	}
}

func listLiquidation(msg string) {
	var pushObj wsPush.LiquidationResp
	err := sJson.ToObj(msg, &pushObj)
	if err != nil {
		sLog.Error("listLiquidation: sJson.ToObj(): %s", err)
		return
	}

	// sLog.Debug("Topic: %s  ..  Type: %s  ..  TimeStamp: %d", pushObj.Topic, pushObj.Type, pushObj.TimeStamp)

	sLog.Info("%s: %+v ", pushObj.Topic, pushObj.Data)
}

func listTickers(msg string) {
	var pushObj wsPush.TickerResp
	err := sJson.ToObj(msg, &pushObj)
	if err != nil {
		sLog.Error("listTickers: sJson.ToObj(): %s", err)
		return
	}

	// sLog.Debug("Topic: %s  ..  Type: %s  ..  TimeStamp: %d", pushObj.Topic, pushObj.Type, pushObj.TimeStamp)

	sLog.Debug("%s: %+v ", pushObj.Topic, pushObj.Data)
}
