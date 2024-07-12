package user

import (
	"github.com/yasseldg/bybit/common"
)

type AffiliateUserInfo struct {
	uid string
}

func NewAffiliateUserInfo(uid string) (*AffiliateUserInfo, error) {
	return &AffiliateUserInfo{
		uid: uid,
	}, nil
}

// Params returns the order parameters
func (o *AffiliateUserInfo) GetParams() common.Params {
	p := common.NewParams()

	p.Set("uid", o.uid)

	return p
}
