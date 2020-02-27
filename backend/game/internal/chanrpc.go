package internal

import (
	"github.com/name5566/leaf/gate"
)

func init() {
	Skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	Skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	agents[a] = nil
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	exitRoom(a)
	if agents[a] != nil {
		agents[a].Agent = nil
	}
	delete(agents, a)
}
