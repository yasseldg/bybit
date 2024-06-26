package response

// UserVolumeInfo represents the structure for a user's trading and deposit volume information.
type UserVolumeInfo struct {
	UID                 string `json:"uid"`
	VipLevel            string `json:"vipLevel"`
	TakerVol30Day       string `json:"takerVol30Day"`
	MakerVol30Day       string `json:"makerVol30Day"`
	TradeVol30Day       string `json:"tradeVol30Day"`
	DepositAmount30Day  string `json:"depositAmount30Day"`
	TakerVol365Day      string `json:"takerVol365Day"`
	MakerVol365Day      string `json:"makerVol365Day"`
	TradeVol365Day      string `json:"tradeVol365Day"`
	DepositAmount365Day string `json:"depositAmount365Day"`
	TotalWalletBalance  string `json:"totalWalletBalance"` // This should be an integer value representing a range, not a string.
	DepositUpdateTime   string `json:"depositUpdateTime"`
	VolUpdateTime       string `json:"volUpdateTime"`
}

// UserAPIKeyInfo represents the structure for a user's API key information.
type UserAPIKeyInfo struct {
	ID            string              `json:"id"`
	Note          string              `json:"note"`
	APIKey        string              `json:"apiKey"`
	ReadOnly      int                 `json:"readOnly"`
	Permissions   map[string][]string `json:"permissions"`
	IPs           []string            `json:"ips"`
	Type          int                 `json:"type"`
	DeadlineDay   int                 `json:"deadlineDay"`
	ExpiredAt     string              `json:"expiredAt"`
	CreatedAt     string              `json:"createdAt"`
	Unified       int                 `json:"unified"`
	UTA           int                 `json:"uta"`
	UserID        int                 `json:"userID"`
	InviterID     int                 `json:"inviterID"`
	VIPLevel      string              `json:"vipLevel"`
	MktMakerLevel string              `json:"mktMakerLevel"`
	AffiliateID   int                 `json:"affiliateID"`
	RSAPublicKey  string              `json:"rsaPublicKey"`
	IsMaster      bool                `json:"isMaster"`
	ParentUID     string              `json:"parentUid"`
	KYCLevel      string              `json:"kycLevel"`
	KYCRegion     string              `json:"kycRegion"`
}
