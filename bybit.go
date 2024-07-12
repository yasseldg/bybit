package bybit

import (
	"github.com/yasseldg/bybit/common"
	"github.com/yasseldg/bybit/rest/service"
)

type InterRest interface {
	NewPlaceOrder(category, symbol, side, orderType string, qty float64, prec int) (*service.PlaceOrder, error)
	NewAmendOrder(category, symbol, orderId string, prec int) (*service.AmendOrder, error)
	NewCancelOrder(category, symbol, orderId string) (*service.CancelOrder, error)
	NewCancelAllOrders(category string) (*service.CancelAllOrders, error)
	NewGetOpenOrders(category string) (*service.GetOpenOrders, error)
	NewGetApiKeyInfo() (*service.GetAPIKeyInfo, error)
	NewGetWalletBalance(accountType string) (*service.GetWalletBalance, error)
	NewGetInstrumentsInfo(category string) (*service.GetInstrumentsInfo, error)
	NewGetAffiliateUserInfo(uid string) (*service.GetAffiliateUserInfo, error)
}

type Rest struct {
	common.InterClient
}

func NewClient(api_key, api_secret string, options ...common.ClientOption) *Rest {
	return &Rest{
		InterClient: common.NewClient(api_key, api_secret, options...),
	}
}

func (c Rest) NewPlaceOrder(category, symbol, side, orderType string, qty float64, prec int) (*service.PlaceOrder, error) {
	return service.NewPlaceOrder(c.InterClient, category, symbol, side, orderType, qty, prec)
}

func (c Rest) NewAmendOrder(category, symbol, orderId string, prec int) (*service.AmendOrder, error) {
	return service.NewAmendOrder(c.InterClient, category, symbol, orderId, prec)
}

func (c Rest) NewCancelOrder(category, symbol, orderId string) (*service.CancelOrder, error) {
	return service.NewCancelOrder(c.InterClient, category, symbol, orderId)
}

func (c Rest) NewGetOpenOrders(category string) (*service.GetOpenOrders, error) {
	return service.NewGetOpenOrders(c.InterClient, category)
}

func (c Rest) NewGetApiKeyInfo() (*service.GetAPIKeyInfo, error) {
	return service.NewGetApiKeyInfo(c.InterClient)
}

func (c Rest) NewGetWalletBalance(accountType string) (*service.GetWalletBalance, error) {
	return service.NewGetWalletBalance(c.InterClient, accountType)
}

func (c Rest) NewCancelAllOrders(category string) (*service.CancelAllOrders, error) {
	return service.NewCancelAllOrders(c.InterClient, category)
}

func (c Rest) NewGetInstrumentsInfo(category string) (*service.GetInstrumentsInfo, error) {
	return service.NewGetInstrumentsInfo(c.InterClient, category)
}

func (c Rest) NewGetAffiliateUserInfo(uid string) (*service.GetAffiliateUserInfo, error) {
	return service.NewGetAffiliateUserInfo(c.InterClient, uid)
}
