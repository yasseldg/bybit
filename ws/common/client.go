package common

import (
	"github.com/yasseldg/bybit/ws/model/wsRequest"

	"github.com/yasseldg/bybit/constants"
)

type WsClient struct {
	BaseWsClient *BaseWsClient
}

func (wsc *WsClient) Init(channel constants.Channel, key, secret string, listener OnReceive, errorListener OnReceive) {
	wsc.BaseWsClient = new(BaseWsClient)
	wsc.BaseWsClient.Init(channel, key, secret)
	wsc.BaseWsClient.SetListener(listener, errorListener)
	wsc.BaseWsClient.ConnectWebSocket()
	wsc.BaseWsClient.StartReadLoop()
	wsc.BaseWsClient.ExecuterPing()
	wsc.BaseWsClient.StartTickerLoop()
	wsc.BaseWsClient.Login()
}

func (wsc *WsClient) UnSubscribe(list []constants.SubscribeTopic) {

	var args []interface{}
	for _, req := range list {
		delete(wsc.BaseWsClient.ListenerMap, req)
		wsc.BaseWsClient.AllSubscribe.Add(req)
		wsc.BaseWsClient.AllSubscribe.Remove(req)
		args = append(args, req)
	}

	wsBaseReq := wsRequest.Base{
		Op:   constants.WsOpUnsubscribe,
		Args: args,
	}

	wsc.BaseWsClient.SendByType(wsBaseReq)
}

func (wsc *WsClient) SubscribeDef(list []constants.SubscribeTopic) {

	var args []interface{}
	for _, req := range list {
		args = append(args, req)
	}
	wsBaseReq := wsRequest.Base{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}
	wsc.BaseWsClient.SendByType(wsBaseReq)
}

func (wsc *WsClient) Subscribe(list []constants.SubscribeTopic, listener OnReceive) {

	var args []interface{}
	for _, req := range list {
		args = append(args, req)

		wsc.BaseWsClient.ListenerMap[req] = listener
		wsc.BaseWsClient.AllSubscribe.Add(req)
	}
	wsBaseReq := wsRequest.Base{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}
	wsc.BaseWsClient.SendByType(wsBaseReq)
}

func (wsc *WsClient) Close() {
	wsc.BaseWsClient.DisconnectWebSocket()
}
