package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds, as requested by Binance.
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func GetCurrentTime() int64 {
	now := time.Now()
	unixNano := now.UnixNano()
	timeStamp := unixNano / int64(time.Millisecond)
	return timeStamp
}

func ValidateParams(params map[string]interface{}) error {
	seenKeys := make(map[string]bool)

	for key, value := range params {
		if key == "" {
			return fmt.Errorf("empty key found in parameters")
		}
		if seenKeys[key] {
			return fmt.Errorf("duplicate key found in parameters: %s", key)
		}
		if value == nil {
			return fmt.Errorf("parameter for key '%s' is nil", key)
		}
		seenKeys[key] = true
	}
	return nil
}
