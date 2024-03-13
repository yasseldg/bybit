package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sJson"

	"github.com/google/uuid"
)

// Config define the config

type config struct {
	ApiKey    string
	ApiSecret string
	BaseUrl   string
}

func getConfig(api_key, api_secret string) config {
	config := config{
		ApiKey:    os.Getenv("BYBIT_API_KEY"),
		ApiSecret: os.Getenv("BYBIT_API_SECRET"),
	}

	if len(api_key) > 0 {
		config.ApiKey = api_key
	}

	if len(api_secret) > 0 {
		config.ApiSecret = api_secret
	}

	return config
}

func (c *config) String() string {
	return fmt.Sprintf("ApiKey: %s .. ApiSecret: %s .. BaseUrl: %s", c.ApiKey, c.ApiSecret, c.BaseUrl)
}

// RestConfig define the rest config

type restConfig struct {
	config
	RecvWindow string
}

// GetRestConfig Get the rest config
// by default it will use the environment variables
// by default base_url = "https://api.bybit.com"
func getRestConfig(api_key, api_secret string) *restConfig {
	config := &restConfig{
		config:     getConfig(api_key, api_secret),
		RecvWindow: os.Getenv("BYBIT_RECV_WINDOW"),
	}
	config.BaseUrl = os.Getenv("BYBIT_BASE_URL_REST")

	if len(config.RecvWindow) == 0 {
		config.RecvWindow = "5000"
	}

	if len(config.BaseUrl) == 0 {
		config.BaseUrl = REST_MAINNET
	}

	return config
}

func (rc *restConfig) Log() {
	sLog.Info("RestConfig: %s .. RecvWindow: %s", rc.config.String(), rc.RecvWindow)
}

// WsConfig define the ws config

type wsConfig struct {
	config
}

func getWsConfig(api_key, api_secret string) *wsConfig {
	config := &wsConfig{
		config: getConfig(api_key, api_secret),
	}
	config.BaseUrl = os.Getenv("BYBIT_BASE_URL_WS")

	if len(config.BaseUrl) == 0 {
		config.BaseUrl = WS_MAINNET
	}

	return config
}

func (wc *wsConfig) Log() {
	sLog.Info("WsConfig: %s", wc.config.String())
}

func (wc *wsConfig) GetAuth() (string, error) {
	// Get current Unix time in milliseconds
	expires := time.Now().UnixNano()/1e6 + 10000
	val := fmt.Sprintf("GET/realtime%d", expires)

	h := hmac.New(sha256.New, []byte(wc.ApiSecret))
	h.Write([]byte(val))

	// Convert to hexadecimal instead of base64
	signature := hex.EncodeToString(h.Sum(nil))
	sLog.Warn("GetAuth signature: %s", signature)

	authMessage := map[string]interface{}{
		"req_id": uuid.New(),
		"op":     "auth",
		"args":   []interface{}{wc.ApiKey, expires, signature},
	}

	return sJson.ToString(authMessage)
}

func GetWsConfig(api_key, api_secret string) *wsConfig {
	return getWsConfig(api_key, api_secret)
}
