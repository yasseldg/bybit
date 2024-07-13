package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/yasseldg/bybit/config"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sJson"

	"github.com/google/uuid"
)

// WsConfig define the ws config

type InterWsConfig interface {
	InterApi

	Log()
	Url() string
	GetAuth() (string, error)
}

type wsConfig struct {
	InterApi

	BaseUrl string
}

func getWsConfig(api_key, api_secret string) *wsConfig {
	ws_config := &wsConfig{
		InterApi: getApi(api_key, api_secret),
	}
	ws_config.BaseUrl = os.Getenv("BYBIT_BASE_URL_WS")

	if len(ws_config.BaseUrl) == 0 {
		ws_config.BaseUrl = string(config.Url_Base_Wss)
	}

	return ws_config
}

func (wc *wsConfig) Log() {
	sLog.Info("WsConfig: %s", wc.InterApi.String())
}

func (wc *wsConfig) Url() string {
	return wc.BaseUrl
}

func (wc *wsConfig) GetAuth() (string, error) {
	// Get current Unix time in milliseconds
	expires := time.Now().UnixNano()/1e6 + 10000
	val := fmt.Sprintf("GET/realtime%d", expires)

	h := hmac.New(sha256.New, []byte(wc.InterApi.Secret()))
	h.Write([]byte(val))

	// Convert to hexadecimal instead of base64
	signature := hex.EncodeToString(h.Sum(nil))

	authMessage := map[string]interface{}{
		"req_id": uuid.New(),
		"op":     "auth",
		"args":   []interface{}{wc.InterApi.Key(), expires, signature},
	}

	return sJson.ToString(authMessage)
}

func GetWsConfig(api_key, api_secret string) *wsConfig {
	return getWsConfig(api_key, api_secret)
}
