package util

import (
	"bytes"
	"fmt"
	"rota/db"
	"strconv"
	"time"

	"github.com/name5566/leaf/log"
)

// IsEmpty check param zero value
func IsEmpty(val interface{}) bool {
	switch v := val.(type) {
	case int:
		return v == 0
	case int32:
		return v == 0
	case int64:
		return v == 0
	case float32:
		return v == 0
	case float64:
		return v == 0
	case string:
		return v == ""
	case []byte:
		return bytes.Equal(v, []byte{0}) || string(v) == ""
	case []string:
		return len(v) == 0
	default:
		s := fmt.Sprintf("%d", val)
		return s == "0"
	}
}

// RequireParam check param and return true for all param provided
func RequireParam(v ...interface{}) bool {
	for _, arg := range v {
		if IsEmpty(arg) {
			return false
		}
	}
	return true
}

// CheckToken check and update token
func CheckToken(token string) (int, string) {
	if token == "" {
		return 0, ""
	}

	db.InitRS() // if opened it will not open again
	res, err := db.RSClient.HMGet(token, "ID", "Name").Result()
	if err != nil || res[0] == nil || res[1] == nil {
		return 0, ""
	}
	db.RSClient.Expire(token, 30*time.Minute) // reset expire time

	id, err := strconv.Atoi(res[0].(string))
	if err != nil {
		log.Error("%v", err)
	}
	// no need to close redis
	return id, res[1].(string)
}
