package msg

import (
	"reflect"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Response{})
	LoginMsgInit()
	RoomMsgInit()
	ChatMsgInit()
}

func Send(a gate.Agent, m interface{}) {
	if a != nil {
		a.WriteMsg(&Response{Type: reflect.ValueOf(m).Elem().Type().String()[4:], Msg: m})
	}
}

func SendError(a gate.Agent, s string) {
	if a != nil {
		a.WriteMsg(&Response{Type: "Error", Msg: &Error{Info: s}})
	}
}

type Response struct {
	Type string
	Msg  interface{}
}

type Error struct {
	Info string
}
