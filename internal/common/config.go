package common

import (
	"os"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type restConfig struct {
	ApiKey     string
	ApiSecret  string
	RecvWindow string
	BaseUrl    string
}

// GetRestConfig Get the rest config
// by default it will use the environment variables
// by default base_url = "https://api.bybit.com"
func getRestConfig(api_key, api_secret string) *restConfig {
	config := &restConfig{
		ApiKey:     os.Getenv("BYBIT_API_KEY"),
		ApiSecret:  os.Getenv("BYBIT_API_SECRET"),
		RecvWindow: os.Getenv("BYBIT_RECV_WINDOW"),
		BaseUrl:    os.Getenv("BYBIT_BASE_URL"),
	}

	if len(api_key) > 0 {
		config.ApiKey = api_key
	}

	if len(api_secret) > 0 {
		config.ApiSecret = api_secret
	}

	if len(config.RecvWindow) == 0 {
		config.RecvWindow = "5000"
	}

	if len(config.BaseUrl) == 0 {
		config.BaseUrl = MAINNET
	}

	return config
}

func (rc *restConfig) Log() {
	sLog.Info("RestConfig: ApiKey: %s .. ApiSecret: %s .. RecvWindow: %s ..  BaseUrl: %s", rc.ApiKey, rc.ApiSecret, rc.RecvWindow, rc.BaseUrl)
}

// DELETE
func init() {
	os.Setenv("BYBIT_API_KEY", "SlX0sUOvRisU6nr3NP")
	os.Setenv("BYBIT_API_SECRET", "yfNquA8D28KIWM16cq2BmtrfGLkVXqg5r2mD")
}
