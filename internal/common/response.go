package common

import (
	"fmt"
)

type RetResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}

type BaseResponse struct {
	RetResponse
	Time int64 `json:"time"`
}

type InfoResponse struct {
	BaseResponse
	RetExtInfo struct{} `json:"retExtInfo"`
}

type ServerResponse struct {
	InfoResponse
	Result interface{} `json:"result"`
}

// Error return error code and message
func (rr RetResponse) Error() error {
	if rr.RetCode == 0 {
		return nil
	}
	return fmt.Errorf("code=%d, msg=%s", rr.RetCode, rr.RetMsg)
}

func (rr RetResponse) APIError() error {
	return fmt.Errorf("<APIError> code=%d, msg=%s", rr.RetCode, rr.RetMsg)
}
