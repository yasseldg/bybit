package affiliate

import (
	"github.com/yasseldg/bybit/common"
)

type UserList struct {
	size   *int
	cursor *string
}

func NewAffiliateUserList() (*UserList, error) {
	return &UserList{}, nil
}

// SetSize sets the size parameter
func (o *UserList) SetSize(size int) {
	o.size = &size
}

// SetCursor sets the cursor parameter
func (o *UserList) SetCursor(cursor string) {
	o.cursor = &cursor
}

// Params returns the order parameters
func (o *UserList) GetParams() common.Params {
	p := common.NewParams()

	if o.size != nil {
		p.Set("size", *o.size)
	}

	if o.cursor != nil {
		p.Set("cursor", *o.cursor)
	}

	return p
}

// Request Parameters

// Parameter	Required	Type	Comments

// size	false	integer	Limit for data size per page. [0, 1000]. Default: 0
// cursor	false	string	Cursor. Use the nextPageCursor token from the response to retrieve the next page of the result set
