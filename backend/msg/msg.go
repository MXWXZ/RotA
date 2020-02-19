package msg

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Response{})
	LoginMsgInit()
	RoomMsgInit()
}

func Send403(a gate.Agent, t string) {
	Send(a, 403, t, "")
}

func Send400(a gate.Agent, t string) {
	Send(a, 400, t, "")
}

func Send500(a gate.Agent, t string, m error) {
	log.Error("%v", m)
	Send(a, 500, t, "")
}

func Send200(a gate.Agent, t string, m interface{}) {
	Send(a, 200, t, m)
}

func Send(a gate.Agent, c uint32, t string, m interface{}) {
	a.WriteMsg(&Response{c, t, m})
}

/**
 * @apiDefine InvalidParam
 * @apiError (Error 400) InvalidParam Invalid param
 */

/**
 * @apiDefine ServerError
 * @apiError (Error 500) ServerError Server error
 */

type Response struct {
	Code uint32
	Type string
	Msg  interface{}
}
