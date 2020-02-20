package util

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"
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

func GetArgs(args []interface{}, m interface{}) (gate.Agent, string) {
	v := reflect.ValueOf(m)
	v.Elem().Set(reflect.ValueOf(args[0]))
	return args[1].(gate.Agent), v.Elem().Type().String()[5:]
}

func Handler(s *module.Skeleton, m interface{}, h interface{}) {
	s.RegisterChanRPC(reflect.TypeOf(m), h)
}
