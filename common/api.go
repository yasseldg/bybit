package common

import (
	"fmt"
	"os"
)

type InterApi interface {
	Key() string
	Secret() string

	String() string
}

type api struct {
	ApiKey    string
	ApiSecret string
}

func (c *api) Key() string {
	return c.ApiKey
}

func (c *api) Secret() string {
	return c.ApiSecret
}

func (c *api) String() string {
	return fmt.Sprintf("ApiKey: %s .. ApiSecret: %s", c.ApiKey, c.ApiSecret)
}

// private funcs

func getApi(api_key, api_secret string) *api {
	api := &api{
		ApiKey:    os.Getenv("BYBIT_API_KEY"),
		ApiSecret: os.Getenv("BYBIT_API_SECRET"),
	}

	if len(api_key) > 0 {
		api.ApiKey = api_key
	}

	if len(api_secret) > 0 {
		api.ApiSecret = api_secret
	}

	return api
}
