package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

type Signer struct {
	secretKey []byte
}

func (s *Signer) Init(key string) *Signer {
	s.secretKey = []byte(key)
	return s
}

func (s *Signer) Sign(method string, requestPath string, body string, timesStamp string) string {

	var payload strings.Builder
	payload.WriteString(timesStamp)
	payload.WriteString(method)
	payload.WriteString(requestPath)
	if body != "" && body != "?" {
		payload.WriteString(body)
	}
	hash := hmac.New(sha256.New, s.secretKey)
	hash.Write([]byte(payload.String()))
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}
