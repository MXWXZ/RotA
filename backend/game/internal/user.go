package internal

import (
	"rota/msg"

	"github.com/name5566/leaf/gate"
)

var agents = make(map[gate.Agent]*User)
var users = make(map[int]*User)

type User struct {
	ID     int // 0 for unknown
	Name   string
	Token  string
	Room   int // 0 for no room
	Status int // 0 for none 1 for gaming
	Agent  gate.Agent
}

func checkUser(f func(args []interface{})) func(args []interface{}) {
	return func(args []interface{}) {
		if v, ok := agents[args[1].(gate.Agent)]; ok && v != nil {
			f(args)
		} else {
			msg.Send(args[1].(gate.Agent), &msg.NeedTokenRsp{})
		}
	}
}

func broadCast(m interface{}) {
	for i := range agents {
		msg.Send(i, m)
	}
}

func broadCastRoom(r int, m interface{}) {
	for i, v := range agents {
		if v.Room == r {
			msg.Send(i, m)
		}
	}
}

func broadCastGroup(u []int, m interface{}) {
	for _, v := range u {
		if _, ok := users[v]; ok && users[v].Agent != nil {
			msg.Send(users[v].Agent, m)
		}
	}
}
