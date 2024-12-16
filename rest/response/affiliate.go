package response

type AffiliateUserList struct {
	List           []UserInfo `json:"list"`
	NextPageCursor string     `json:"nextPageCursor"`
}

type UserInfo struct {
	UserId       string `json:"userId"`
	RegisterTime string `json:"registerTime"`
	Source       string `json:"source"`
	Remarks      string `json:"remarks"`
	IsKyc        bool   `json:"isKyc"`
}

// Response Parameters

// Parameter	Type	Comments

// list	array	Object
// > userId	string	user Id
// > registerTime	string	user register time
// > source	integer	user registration source, from which referrer code
// > remarks	integer	The remark
// > isKyc	boolean	Whether KYC is completed
// nextPageCursor	string	Refer to the cursor request parameter
