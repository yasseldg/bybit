package common

import (
	"os"

	"github.com/yasseldg/bybit/config"

	"github.com/yasseldg/go-simple/logs/sLog"
)

// RestConfig define the rest config

type InterRestConfig interface {
	InterApi

	Log()
	Url() string
}

type restConfig struct {
	InterApi

	BaseUrl    string
	RecvWindow string
}

func (rc *restConfig) Url() string {
	return rc.BaseUrl
}

// GetRestConfig Get the rest config
// by default it will use the environment variables
// by default base_url = "https://api.bybit.com"
func getRestConfig(api_key, api_secret string) *restConfig {
	rest_config := &restConfig{
		InterApi:   getApi(api_key, api_secret),
		RecvWindow: os.Getenv("BYBIT_RECV_WINDOW"),
	}
	rest_config.BaseUrl = os.Getenv("BYBIT_BASE_URL_REST")

	if len(rest_config.RecvWindow) == 0 {
		rest_config.RecvWindow = "5000"
	}

	if len(rest_config.BaseUrl) == 0 {
		rest_config.BaseUrl = string(config.Url_Base_Rest)
	}

	return rest_config
}

func (rc *restConfig) Log() {
	sLog.Info("RestConfig: %s .. RecvWindow: %s", rc.InterApi.String(), rc.RecvWindow)
}
