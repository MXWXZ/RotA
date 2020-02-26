package internal

import (
	"rota/msg"

	"github.com/name5566/leaf/gate"
)

var agents = make(map[gate.Agent]*User)

type User struct {
	ID     int // 0 for unknown
	Name   string
	Room   int // 0 for no room
	Status int // 0 for none 1 for gaming
}

func (u *User) Reset() {
	u.ID = 0
	u.Name = ""
	u.Room = 0
	u.Status = 0
}

func checkUser(f func(args []interface{})) func(args []interface{}) {
	return func(args []interface{}) {
		if v, ok := agents[args[1].(gate.Agent)]; ok && v.ID != 0 {
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
