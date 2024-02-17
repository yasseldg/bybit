package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type secType int

const (
	secTypeNone   secType = iota
	secTypeSigned         // private Request
)

// Request define an API Request
type Request struct {
	method     string
	endpoint   string
	query      url.Values
	recvWindow string
	secType    secType
	header     http.Header
	params     []byte
	fullURL    string
	body       io.Reader
}

func (r *Request) validate() (err error) {
	if r.query == nil {
		r.query = url.Values{}
	}
	return nil
}

func (r *Request) Get() *Request {
	r.method = http.MethodGet
	return r
}

func (r *Request) Post() *Request {
	r.method = http.MethodPost
	return r
}

func (r *Request) Signed() *Request {
	r.secType = secTypeSigned
	return r
}

func (r *Request) EndPoint(endpoint string) *Request {
	r.endpoint = endpoint
	return r
}

func (r *Request) SetParams(params url.Values) error {
	r.query = params

	switch r.method {
	case http.MethodGet:

	case http.MethodPost:
		jsonData, err := json.Marshal(r.query)
		if err != nil {
			return fmt.Errorf("error marshalling query: %s", err)
		}
		r.params = jsonData
	}

	return nil
}
