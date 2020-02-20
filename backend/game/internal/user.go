package internal

import (
	"github.com/name5566/leaf/gate"
)

var agents = make(map[gate.Agent]*User)

type User struct {
	ID   int // 0 for unknown
	Name string
	Room int // 0 for no room
}

func (u *User) Reset() {
	u.ID = 0
	u.Name = ""
	u.Room = 0
}

func checkUser(f func(args []interface{})) func(args []interface{}) {
	return func(args []interface{}) {
		if v, ok := agents[args[1].(gate.Agent)]; ok && v.ID != 0 {
			f(args)
		}
	}
}
