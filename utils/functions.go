package utils

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/yasseldg/bybit/constants"
)

func TimesStamp() string {
	timesStamp := time.Now().Unix() * 1000
	return strconv.FormatInt(timesStamp, 10)
}

func TimesStampSec() string {
	timesStamp := time.Now().Unix()
	return strconv.FormatInt(timesStamp, 10)
}

/**
 * get header
 */
func Headers(request *http.Request, apikey string, timestamp string, sign string, passphrase string) {
	request.Header.Add(constants.ContentType, constants.ApplicationJson)
	request.Header.Add(constants.AccessKey, apikey)
	request.Header.Add(constants.AccessSign, sign)
	request.Header.Add(constants.AccessTimestamp, timestamp)
	request.Header.Add(constants.AccessPassphrase, passphrase)
}

func BuildJsonParams(params map[string]string) (string, error) {
	if params == nil {
		return "", errors.New("illegal parameter")
	}
	data, err := json.Marshal(params)
	if err != nil {
		return "", errors.New("json convert string error")
	}
	jsonBody := string(data)
	return jsonBody, nil
}

func BuildGetParams(params map[string]string) string {
	urlParams := url.Values{}
	if len(params) > 0 {
		for k := range params {
			urlParams.Add(k, params[k])
		}
	}
	return "?" + urlParams.Encode()
}

func NewParams() map[string]string {
	return make(map[string]string)
}

func powerf(x float64, n int) float64 {
	ans := 1.0
	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}

func GetSignedInt(checksum string) string {
	c, _ := strconv.ParseUint(checksum, 10, 64)

	if c > math.MaxInt32 {
		a := c - (1<<31-1)*2 - 2
		return strconv.FormatUint(a, 10)
	}
	return checksum
}
