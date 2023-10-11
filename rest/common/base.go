package common

import (
	"encoding/json"
	"fmt"

	"github.com/yasseldg/bybit/constants"
	"github.com/yasseldg/simplego/sNet"
)

type BaseRestClient struct {
	Service     sNet.Config
	NeedLogin   bool
	Connection  bool
	LoginStatus bool
}

func (b *BaseRestClient) Init(env string, needLogin bool) {
	b.NeedLogin = needLogin

	if len(env) == 0 {
		env = constants.Env_DEFAULT
	}

	s := sNet.GetConfig(env)
	if s == nil {
		b.Connection = false
		return
	}
	s.Log()

	b.Service = *s
	b.Connection = true
}

func (b *BaseRestClient) Login() {
	b.LoginStatus = true
}

func (rc *BaseRestClient) Call(params string, response interface{}) error {

	body, err := rc.Service.Call(constants.GET, params, nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("json.Unmarshal(body, response): %s ", err)
	}

	return nil
}
