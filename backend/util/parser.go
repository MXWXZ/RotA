package util

import (
	"bytes"
	"fmt"
	"runtime"
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

func GetFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
