package gate

import (
	"rota/game"
	"rota/login"
	"rota/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.Signup{}, login.ChanRPC)

	msg.Processor.SetRouter(&msg.CheckToken{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.GetRooms{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.NewRoom{}, game.ChanRPC)
}
