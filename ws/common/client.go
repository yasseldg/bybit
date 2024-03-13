package common

import (
	"time"

	"github.com/yasseldg/bybit/ws/model/wsRequest"

	"github.com/yasseldg/bybit/constants"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type WsClient struct {
	BaseWsClient *BaseWsClient
}

func (wsc *WsClient) Init(channel constants.Channel, needLogin bool, listener OnReceive, errorListener OnReceive) {
	wsc.BaseWsClient = new(BaseWsClient)
	wsc.BaseWsClient.Init(channel, needLogin)
	wsc.BaseWsClient.SetListener(listener, errorListener)
	wsc.BaseWsClient.ConnectWebSocket()
	wsc.BaseWsClient.StartReadLoop()
	wsc.BaseWsClient.ExecuterPing()
	wsc.BaseWsClient.StartTickerLoop()

	// TO FIX
	if needLogin {
		sLog.Info("WebSocket login in ...")
		c := 0
		for {
			wsc.BaseWsClient.Login()
			time.Sleep(1 * time.Second)
			c++

			if c > 10 {
				sLog.Error("WebSocket login in ... failed")
				break
			}

			if !wsc.BaseWsClient.LoginStatus {
				continue
			}
			break
		}
		sLog.Info("WebSocket login in ... success")
	}
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

	wsc.SendMessageByType(wsBaseReq)
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
	wsc.SendMessageByType(wsBaseReq)
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
	wsc.SendMessageByType(wsBaseReq)
}

func (wsc *WsClient) Connect() {
	wsc.BaseWsClient.Connect()
}

func (wsc *WsClient) SendMessage(msg string) {
	wsc.BaseWsClient.Send(msg, false)
}

func (wsc *WsClient) SendMessageByType(req wsRequest.Base) {
	wsc.BaseWsClient.SendByType(req)
}

func (wsc *WsClient) Close() {
	wsc.BaseWsClient.DisconnectWebSocket()
}
