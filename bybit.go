package bybit

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/service"
)

type Rest struct {
	*common.Client
}

func NewClient(api_key, api_secret string, options ...common.ClientOption) *Rest {
	return &Rest{
		Client: common.NewClient(api_key, api_secret, options...),
	}
}

func NewPlaceOrder(c *common.Client, category, symbol, side, orderType string, qty float64, prec int) (*service.PlaceOrder, error) {
	return service.NewPlaceOrder(c, category, symbol, side, orderType, qty, prec)
}

func NewAmendOrder(c *common.Client, category, symbol, orderId string, prec int) (*service.AmendOrder, error) {
	return service.NewAmendOrder(c, category, symbol, orderId, prec)
}

func NewCancelOrder(c *common.Client, category, symbol, orderId string) (*service.CancelOrder, error) {
	return service.NewCancelOrder(c, category, symbol, orderId)
}

func NewGetOpenOrders(c *common.Client, category string) (*service.GetOpenOrders, error) {
	return service.NewGetOpenOrders(c, category)
}

func NewGetApiKeyInfo(c *common.Client) (*service.GetAPIKeyInfo, error) {
	return service.NewGetApiKeyInfo(c)
}

func NewGetWalletBalance(c *common.Client, accountType string) (*service.GetWalletBalance, error) {
	return service.NewGetWalletBalance(c, accountType)
}
