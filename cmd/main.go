package main

import (
	"context"
	"os"
	"time"

	"github.com/yasseldg/bybit"
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/constants"
	"github.com/yasseldg/bybit/ws/model/wsPush"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sTime"

	"github.com/yasseldg/simplego/sJson"
)

func main() {
	clean := sLog.SetByName(sLog.Zap, sLog.LevelInfo, "")
	defer clean()

	sLog.Info("Starting...")

	sTime.TimeControl(rest, "Start")
}

func rest() {

	os.Setenv("Serv_Bybit_API", "API_v5")

	rest := bybit.NewClient("", "", common.WithDebug(true))
	rest.Log()

	restInstrumentsInfos(rest)
	println()

	restGetApiKeyInfo(rest)
	println()

	restWalletBalance(rest)
	println()
}

func restWalletBalance(rest bybit.InterRest) {

	wallet, err := rest.NewGetWalletBalance("UNIFIED")
	if err != nil {
		sLog.Error("bybit.NewGetWalletBalance(): %s", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	defer cancel()

	resp, err := wallet.Do(ctx)
	if err != nil {
		sLog.Error("wallet.Do(): %s", err)
		return
	}

	sLog.Info("WalletBalance: %+v", resp.Result)
}

func restGetApiKeyInfo(rest bybit.InterRest) {

	info, err := rest.NewGetApiKeyInfo()
	if err != nil {
		sLog.Error("bybit.NewGetApiKeyInfo(): %s", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	defer cancel()

	resp, err := info.Do(ctx)
	if err != nil {
		sLog.Error("info.Do(): %s", err)
		return
	}

	sLog.Info("GetApiKeyInfo: %+v", resp.Result)
}

func restInstrumentsInfos(rest bybit.InterRest) {

	instruments, err := rest.NewGetInstrumentsInfo(string(constants.Category_Linear))
	if err != nil {
		sLog.Error("bybit.NewGetInstrumentsInfo(): %s", err)
		return
	}

	instruments.Symbol("BTCUSDT")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)
	defer cancel()

	resp, err := instruments.Do(ctx)
	if err != nil {
		sLog.Error("instruments.Do(): %s", err)
		return
	}

	sLog.Info("InstrumentsInfos: %+v", resp.Result)
}

func wss() {
	wssLinear() //  comment for test Inverse
	wssInverse()
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
