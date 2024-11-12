package common

import (
	"encoding/json"
	"time"

	"github.com/yasseldg/bybit/constants"
	"github.com/yasseldg/go-simple/logs/sLog"
)

// Bucle de lectura del WebSocket
func (p *BaseWsClient) readLoop() {
	for {
		if p.WebSocketClient == nil {
			sLog.Warn("WebSocket Read error: no connection available")
			time.Sleep(time.Second)
			continue
		}

		_, buf, err := p.WebSocketClient.ReadMessage()
		if err != nil {
			sLog.Warn("WebSocket Read error: %s", err)
			time.Sleep(time.Second)
			continue
		}

		p.LastReceivedTime = time.Now()

		// Enviar el mensaje al canal para ser procesado por un worker
		p.messageChannel <- buf
	}
}

// Worker para procesar mensajes de forma concurrente
func (p *BaseWsClient) worker() {
	for message := range p.messageChannel {
		p.processMessage(message)
	}
}

func (p *BaseWsClient) processMessage(buf []byte) {
	var jsonMap map[string]interface{}

	err := json.Unmarshal(buf, &jsonMap)
	if err != nil {
		sLog.Warn("WebSocket Read error: json.Unmarshal: %s", err)
		return
	}

	v, e := jsonMap["op"]
	if e && v.(string) == constants.WsOpPong {
		sLog.Debug("WebSocket keep connected %s", constants.WsOpPong)
		return
	}

	if e && v.(string) == constants.WsOpPing {
		v, e := jsonMap["success"]
		if e && v.(bool) {
			sLog.Debug("WebSocket keep connected %s", constants.WsOpPing)
			return
		}
	}

	if e && v.(string) == constants.WsOpAuth {
		v, e := jsonMap["success"]
		if e && v.(bool) {
			sLog.Info("WebSocket login success")
			p.LoginStatus = true
			return
		}

		msg := ""
		v, e = jsonMap["ret_msg"]
		if e {
			msg = v.(string)
		}
		sLog.Error("WebSocket login failed: %s", msg)
		return
	}

	message := string(buf)

	v, e = jsonMap["topic"]
	if e {
		listener := p.GetListener(constants.SubscribeTopic(v.(string)))

		listener(message)
		return
	}
	p.handleMessage(message)
}
