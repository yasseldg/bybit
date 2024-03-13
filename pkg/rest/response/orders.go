package response

import (
	"github.com/yasseldg/bybit/pkg/common"

	internalCommon "github.com/yasseldg/bybit/internal/common"
)

type OrderResult struct {
	OrderId     string `json:"orderId"`
	OrderLinkId string `json:"orderLinkId"`
}

type ListOrderResult struct {
	List []OrderResult `json:"list"`
}

type OrderInfo struct {
	common.Order

	IsLeverage     string `json:"isLeverage"`
	OcoTriggerType string `json:"ocoTriggerType"`
}

type OpenOrdersInfo struct {
	Category       string      `json:"category"`
	NextPageCursor string      `json:"nextPageCursor"`
	List           []OrderInfo `json:"list"`
}

type BorrowQuotaInfo struct {
	Symbol             string `json:"symbol"`
	Side               string `json:"side"`
	MaxTradeQty        string `json:"maxTradeQty"`
	MaxTradeAmount     string `json:"maxTradeAmount"`
	SpotMaxTradeQty    string `json:"spotMaxTradeQty"`
	SpotMaxTradeAmount string `json:"spotMaxTradeAmount"`
	BorrowCoin         string `json:"borrowCoin"`
}

type BatchOrderServerResponse struct {
	internalCommon.BaseResponse
	Result struct {
		List []struct {
			Category    string  `json:"category"`
			Symbol      string  `json:"symbol"`
			OrderId     string  `json:"orderId"`
			OrderLinkId string  `json:"orderLinkId"`
			CreateAt    *string `json:"createAt,omitempty"`
		} `json:"list"`
	} `json:"result"`
	RetExtInfo struct {
		List []struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"list"`
	} `json:"retExtInfo"`
}
