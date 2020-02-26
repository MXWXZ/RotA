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
}

func Send(a gate.Agent, m interface{}) {
	a.WriteMsg(&Response{Type: reflect.ValueOf(m).Elem().Type().String()[4:], Msg: m})
}

type Response struct {
	Type string
	Msg  interface{}
}
