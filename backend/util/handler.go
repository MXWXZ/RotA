package util

import (
	"reflect"
	"rota/conf"
	"rota/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/module"
)

func GetArgs(args []interface{}, m interface{}) gate.Agent {
	reflect.ValueOf(m).Elem().Set(reflect.ValueOf(args[0]))
	return args[1].(gate.Agent)
}

func Handler(s *module.Skeleton, m interface{}, h interface{}) {
	s.RegisterChanRPC(reflect.TypeOf(m), h)
}

func SumSlice(s []msg.RoomMember) [conf.MaxRoomTeam]int {
	var ret [conf.MaxRoomTeam]int
	for _, v := range s {
		ret[v.Team] += 1
	}
	return ret
}
