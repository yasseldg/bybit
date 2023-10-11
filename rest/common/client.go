package common

import (
	"fmt"
	"path"
	"time"

	"github.com/yasseldg/bybit/constants"
	"github.com/yasseldg/simplego/sLog"
)

type RestClient struct {
	BaseRestClient *BaseRestClient
	Recourse       constants.Recourse
	EndPoint       constants.EndPoint
}

func (rc *RestClient) Init(env string, needLogin bool) {
	rc.BaseRestClient = new(BaseRestClient)
	rc.BaseRestClient.Init(env, needLogin)
	rc.Login()
}

func (rc *RestClient) Login() {
	// TO FIX
	if rc.BaseRestClient.NeedLogin {
		sLog.Info("Rest login in ...")
		c := 0
		for {
			rc.BaseRestClient.Login()
			time.Sleep(1 * time.Second)
			c++

			if c > 10 {
				sLog.Error("Rest login in ... failed")
				break
			}

			if !rc.BaseRestClient.LoginStatus {
				continue
			}
			break
		}
		sLog.Info("Rest login in ... success")
	}
}

func (rc *RestClient) SetRecourse(recourse constants.Recourse) *RestClient {
	rc.Recourse = recourse
	return rc
}

func (rc *RestClient) SetEndPoint(end_point constants.EndPoint) *RestClient {
	rc.EndPoint = end_point
	return rc
}

func (rc *RestClient) Call(params string, response interface{}) error {
	params = fmt.Sprintf("%s?%s", path.Join(string(rc.Recourse), string(rc.EndPoint)), params)
	return rc.BaseRestClient.Call(params, response)
}
