package common

type ClientOption func(*Client)

// WithDebug print more details in debug mode
func WithDebug(debug bool) ClientOption {
	return func(c *Client) {
		c.Debug = debug
	}
}

// WithBaseUrl is a client option to set the base URL of the Bybit HTTP client.
func WithBaseUrl(base_url string) ClientOption {
	return func(c *Client) {
		c.BaseUrl = base_url
	}
}

// RequestOption define option type for Request
type RequestOption func(*Request)

// WithRecvWindow Append `WithRecvWindow(insert_recvWindow)` to Request to modify the default recvWindow value
func WithRecvWindow(recvWindow string) RequestOption {
	return func(r *Request) {
		r.recvWindow = recvWindow
	}
}
