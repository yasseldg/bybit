package model

import (
	"fmt"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type Order struct {
	OrderId            string `json:"orderId"`
	Side               string `json:"side"`
	Symbol             string `json:"symbol"`
	OrderType          string `json:"orderType"`
	CancelType         string `json:"cancelType"`
	OrderStatus        string `json:"orderStatus"`
	OrderIv            string `json:"orderIv"`
	Price              string `json:"price"`
	Qty                string `json:"qty"`
	TimeInForce        string `json:"timeInForce"`
	OrderLinkId        string `json:"orderLinkId"`
	BlockTradeId       string `json:"blockTradeId"`
	PositionIdx        int    `json:"positionIdx"`
	RejectReason       string `json:"rejectReason"`
	AvgPrice           string `json:"avgPrice"`
	LeavesQty          string `json:"leavesQty"`
	LeavesValue        string `json:"leavesValue"`
	CumExecQty         string `json:"cumExecQty"`
	CumExecValue       string `json:"cumExecValue"`
	CumExecFee         string `json:"cumExecFee"`
	StopOrderType      string `json:"stopOrderType"`
	TriggerPrice       string `json:"triggerPrice"`
	TakeProfit         string `json:"takeProfit"`
	StopLoss           string `json:"stopLoss"`
	TpslMode           string `json:"tpslMode"`
	TpLimitPrice       string `json:"tpLimitPrice"`
	SlLimitPrice       string `json:"slLimitPrice"`
	TpTriggerBy        string `json:"tpTriggerBy"`
	SlTriggerBy        string `json:"slTriggerBy"`
	TriggerDirection   int    `json:"triggerDirection"`
	TriggerBy          string `json:"triggerBy"`
	LastPriceOnCreated string `json:"lastPriceOnCreated"`
	ReduceOnly         bool   `json:"reduceOnly"`
	CloseOnTrigger     bool   `json:"closeOnTrigger"`
	PlaceType          string `json:"placeType"`
	SmpType            string `json:"smpType"`
	SmpGroup           int    `json:"smpGroup"`
	SmpOrderId         string `json:"smpOrderId"`
	CreatedTime        string `json:"createdTime"`
	UpdatedTime        string `json:"updatedTime"`
}

func (o *Order) String() string {
	return fmt.Sprintf("%s .. %s .. %s .. %s .. %s .. %s", o.OrderId, o.OrderLinkId, o.Symbol, o.Side, o.StopOrderType, o.OrderStatus)
}

func (o *Order) Details(name string) string {
	return fmt.Sprintf("Order ( %s ): id: %s .. link_id: %s .. %s .. %s .. %s .. type: %s .. stop_type: %s .. pos_idx: %d .. dir: %d .. qty: %s .. trig_price: %s .. c_at: %s, u_at: %s",
		name, o.OrderId, o.OrderLinkId, o.Symbol, o.Side, o.OrderStatus, o.OrderType, o.StopOrderType, o.PositionIdx, o.TriggerDirection, o.Qty, o.TriggerPrice, o.CreatedTime, o.UpdatedTime)
}

func (o *Order) Log(name string) {
	sLog.Info(o.Details(name))
}
