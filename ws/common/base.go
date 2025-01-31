package common

import (
	"fmt"
	"sync"
	"time"

	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/constants"
	"github.com/yasseldg/bybit/ws/model"
	"github.com/yasseldg/bybit/ws/model/wsRequest"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sJson"

	"github.com/gorilla/websocket"
	"github.com/robfig/cron"
)

type BaseWsClient struct {
	config common.InterWsConfig

	NeedLogin        bool
	Connection       bool
	LoginStatus      bool
	Channel          constants.Channel
	Listener         OnReceive
	ErrorListener    OnReceive
	Ticker           *time.Ticker
	SendMutex        *sync.Mutex
	WebSocketClient  *websocket.Conn
	LastReceivedTime time.Time
	AllSubscribe     *model.Set
	Signer           *Signer
	ListenerMap      map[constants.SubscribeTopic]OnReceive

	workers *workerPool

	messageChannel chan []byte
}

type OnReceive func(message string)

func (p *BaseWsClient) Init(channel constants.Channel, key, secret string) {
	p.config = common.GetWsConfig(key, secret)

	if len(p.config.Key()) > 0 {
		p.NeedLogin = true
	}

	p.Channel = channel
	p.Connection = false
	p.AllSubscribe = model.NewSet()
	p.Signer = new(Signer).Init(p.config.Secret())
	p.ListenerMap = make(map[constants.SubscribeTopic]OnReceive)
	p.SendMutex = &sync.Mutex{}
	p.LastReceivedTime = time.Now()

	p.messageChannel = make(chan []byte, 1000)

	p.workers = newWorkerPool(40, p.messageChannel, p.processMessage)

	if constants.TimerIntervalSecond > 0 {
		p.Ticker = time.NewTicker(constants.TimerIntervalSecond * time.Second)
	}
}

func (p *BaseWsClient) SetWorkers(count int) {
	p.workers.Scale(count)
}

func (p *BaseWsClient) SetListener(msgListener OnReceive, errorListener OnReceive) {
	p.Listener = msgListener
	p.ErrorListener = errorListener
}

func (p *BaseWsClient) Connect() {
	p.ExecuterPing()
	go p.tickerLoop()
}

func (p *BaseWsClient) ConnectWebSocket() {
	var err error

	url := fmt.Sprintf("%s/%s", p.config.Url(), p.Channel)

	sLog.Info("WebSocket connecting ... %s", url)

	p.WebSocketClient, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		sLog.Error("WebSocket connected error: %s\n", err)
		return
	}

	sLog.Info("WebSocket connected: %s", url)
	p.Connection = true
}

// real authentication, Signer is not used
func (p *BaseWsClient) Login() {
	if !p.NeedLogin {
		return
	}

	sLog.Info("WebSocket login in ...")
	c := 0
	for {
		p.login()
		time.Sleep(1 * time.Second)

		if p.LoginStatus {
			sLog.Info("WebSocket login in ... success")
			break
		}

		if c > 10 {
			sLog.Error("WebSocket login in ... failed")
			break
		}
		c++
	}
}

func (p *BaseWsClient) login() {
	auth, err := p.config.GetAuth()
	if err != nil {
		sLog.Error("WebSocket login error: %s", err)
		return
	}
	p.Send(auth, true)
}

// Método para inicializar el cliente y los workers
func (p *BaseWsClient) StartReadLoop() {
	go p.readLoop() // Ejecutar el bucle de lectura en una goroutine

	p.workers.Start() // Iniciar los workers
}

func (p *BaseWsClient) StartTickerLoop() {
	if constants.TimerIntervalSecond > 0 {
		go p.tickerLoop()
	}
}

func (p *BaseWsClient) ExecuterPing() {
	c := cron.New()
	_ = c.AddFunc(fmt.Sprintf("*/%d * * * * *", constants.PingIntervalSecond), p.ping)
	c.Start()
}

func (p *BaseWsClient) ping() {
	json, _ := sJson.ToString(wsRequest.Base{Op: constants.WsOpPing})
	p.Send(json, true)
}

func (p *BaseWsClient) SendByType(req wsRequest.Base) {
	json, _ := sJson.ToString(req)
	p.Send(json, false)
}

func (p *BaseWsClient) Send(data string, debug bool) {
	if p.WebSocketClient == nil {
		sLog.Error("WebSocket sent error: no connection available")
		return
	}

	if debug {
		sLog.Debug("WebSocket sendMessage: %s", data)
	} else {
		sLog.Info("WebSocket sendMessage: %s", data)
	}

	p.SendMutex.Lock()
	err := p.WebSocketClient.WriteMessage(websocket.TextMessage, []byte(data))
	p.SendMutex.Unlock()
	if err != nil {
		sLog.Error("WebSocket sent error: data=%s, error=%s", data, err)
	}
}

func (p *BaseWsClient) SubscribeAll() {

	var args []interface{}
	for _, reg := range p.AllSubscribe.List() {
		args = append(args, reg)
	}

	wsBaseReq := wsRequest.Base{
		Op:   constants.WsOpSubscribe,
		Args: args,
	}
	p.SendByType(wsBaseReq)
}

func (p *BaseWsClient) tickerLoop() {
	sLog.Info("WebSocket TickerLoop started")
	for {
		<-p.Ticker.C
		sLog.Debug("WebSocket ticker")
		elapsedSecond := time.Since(p.LastReceivedTime).Seconds()
		if elapsedSecond > constants.PingIntervalSecond {
			sLog.Info("WebSocket reconnect...")
			p.DisconnectWebSocket()
			p.ConnectWebSocket()
			if p.Connection {
				p.LastReceivedTime = time.Now()

				// Login again
				p.Login()

				// Subscribe All again
				p.SubscribeAll()
			}
		}
	}
}

func (p *BaseWsClient) DisconnectWebSocket() {
	if p.WebSocketClient == nil {
		return
	}

	sLog.Info("WebSocket disconnecting...")

	err := p.WebSocketClient.Close()
	if err != nil {
		sLog.Error("WebSocket disconnect error: %s\n", err)
		return
	}

	sLog.Info("WebSocket disconnected")
}

func (p *BaseWsClient) GetListener(topic constants.SubscribeTopic) OnReceive {
	v, e := p.ListenerMap[topic]
	if !e {
		return p.Listener
	}
	return v
}

func (p *BaseWsClient) handleMessage(msg string) {
	sLog.Info("WebSocket default: %s ", msg)
}
