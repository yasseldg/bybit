package bybit

import (
	"fmt"

	"github.com/yasseldg/bybit/constants"
	"github.com/yasseldg/bybit/rest/common"
	restResponse "github.com/yasseldg/bybit/rest/model/response"
	"github.com/yasseldg/simplego/sLog"
)

// ws
type BybitRestClient struct {
	RestClient *common.RestClient
}

func NewRestClient(env string, needLogin bool) *BybitRestClient {
	rc := new(common.RestClient)
	rc.Init(env, needLogin)
	return &BybitRestClient{rc}
}

func (b *BybitRestClient) SetRecourse(recourse constants.Recourse) *BybitRestClient {
	b.RestClient.SetRecourse(recourse)
	return b
}

func (b *BybitRestClient) SetEndPoint(end_point constants.EndPoint) *BybitRestClient {
	b.RestClient.SetEndPoint(end_point)
	return b
}

func (b BybitRestClient) InstrumentsInfos(category constants.Category, status constants.InstrumentStatus, symbol string) (restResponse.InstrumentsInfos, error) {
	if len(category) == 0 {
		return nil, fmt.Errorf("categorys is empty")
	}

	// ?category=linear,inverse&symbol=BTCUSDT&status=Trading"
	params := fmt.Sprintf("category=%s", category)

	if len(symbol) > 0 {
		params = fmt.Sprintf("%s&symbol=%s", params, symbol)
	}

	if len(status) > 0 {
		params = fmt.Sprintf("%s&status=%s", params, status)
	}

	b.SetRecourse(constants.Recourse_Market)
	b.SetEndPoint(constants.EndPoint_InstrumentsInfo)

	var response restResponse.InstrumentsInfoResp
	err := b.RestClient.Call(params, &response)
	if err != nil {
		return nil, err
	}

	sLog.Debug("InstrumentsInfos: %+v", response)

	return response.Result.List, nil
}

func (b BybitRestClient) Tickers(category constants.Category, symbol string) (restResponse.Tickerss, error) {
	if len(category) == 0 {
		return nil, fmt.Errorf("categorys is empty")
	}

	// ?category=linear,inverse&symbol=BTCUSDT&status=Trading"
	params := fmt.Sprintf("category=%s", category)

	if len(symbol) > 0 {
		params = fmt.Sprintf("%s&symbol=%s", params, symbol)
	}

	b.SetRecourse(constants.Recourse_Market)
	b.SetEndPoint(constants.EndPoint_Tickers)

	var response restResponse.TickersResp
	err := b.RestClient.Call(params, &response)
	if err != nil {
		return nil, err
	}

	sLog.Debug("Tickers: %+v", response)

	return response.Result.List, nil
}
