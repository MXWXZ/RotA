package internal

import (
	"github.com/name5566/leaf/gate"
)

func init() {
	Skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	Skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

var Agents = make(map[gate.Agent]*User)

type User struct {
	ID    int // 0 for unknown
	Name  string
	Room  int // 0 for no room
	Agent gate.Agent
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	Agents[a] = &User{ID: 0}
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	delete(Agents, a)
}
