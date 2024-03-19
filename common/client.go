package common

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/yasseldg/bybit/utils"
)

// Client define API client
type Client struct {
	*restConfig

	HttpClient *http.Client
	Debug      bool
	Logger     *log.Logger
	do         doFunc
}

type doFunc func(req *http.Request) (*http.Response, error)

// NewClient Create client function for initialising new Bybit client
func NewClient(api_key, api_secret string, options ...ClientOption) *Client {
	c := &Client{
		restConfig: getRestConfig(api_key, api_secret),
		HttpClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, Name, log.LstdFlags),
	}

	// Apply the provided options
	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *Client) Log() {
	c.restConfig.Log()
}

func (c *Client) Call(ctx context.Context, r *Request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HttpClient.Do
	}
	res, err := f(req)
	if err != nil {
		return nil, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the returned error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(RetResponse)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr.APIError()
	}
	return data, nil
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

func (c *Client) parseRequest(r *Request, opts ...RequestOption) (err error) {
	// set Request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseUrl, r.endpoint)

	queryString := r.query.Encode()
	header := http.Header{}
	body := &bytes.Buffer{}
	if r.params != nil {
		body = bytes.NewBuffer(r.params)
	}
	if r.header != nil {
		header = r.header.Clone()
	}
	header.Set("User-Agent", fmt.Sprintf("%s/%s", Name, Version))

	if r.secType == secTypeSigned {
		timeStamp := utils.GetCurrentTime()
		header.Set(signTypeKey, "2")
		header.Set(apiRequestKey, c.ApiKey)
		header.Set(timestampKey, strconv.FormatInt(timeStamp, 10))
		if r.recvWindow == "" {
			r.recvWindow = recvWindow
		}
		header.Set(recvWindowKey, r.recvWindow)

		var signatureBase []byte
		if r.method == "POST" {
			header.Set("Content-Type", "application/json")
			signatureBase = []byte(strconv.FormatInt(timeStamp, 10) + c.ApiKey + r.recvWindow + string(r.params[:]))
		} else {
			signatureBase = []byte(strconv.FormatInt(timeStamp, 10) + c.ApiKey + r.recvWindow + queryString)
		}
		hmac256 := hmac.New(sha256.New, []byte(c.ApiSecret))
		hmac256.Write(signatureBase)
		signature := hex.EncodeToString(hmac256.Sum(nil))
		header.Set(signatureKey, signature)
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, body)
	r.fullURL = fullURL
	r.body = body
	r.header = header
	return nil
}
