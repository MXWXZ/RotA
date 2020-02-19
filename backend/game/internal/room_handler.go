package internal

import (
	"reflect"
	"rota/msg"
	"rota/util"

	"github.com/name5566/leaf/gate"
)

func handler(m interface{}, h interface{}) {
	Skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.GetRooms{}, handleGetRooms)
}

var rooms []msg.RoomInfo

func handleGetRooms(args []interface{}) {
	m := args[0].(*msg.GetRooms)
	a := args[1].(gate.Agent)
	t := "GetRooms"

	id, name := util.CheckToken(m.Token)
	if id == 0 {
		msg.Send403(a, t)
		return
	}
	Agents[a].ID = id
	Agents[a].Name = name

	if m.Limit < 0 || m.Offset < 0 {
		msg.Send400(a, t)
		return
	}
	if m.Limit == 0 {
		m.Limit = 20
	} else if m.Limit > 100 {
		m.Limit = 100
	}
	if m.Offset >= int32(len(rooms)) {
		msg.Send200(a, t, &msg.GetRoomsRsp{RoomInfo: make([]msg.RoomInfo, 0, 0)})
		return
	}
	if m.Offset+m.Limit > int32(len(rooms)) {
		m.Limit = int32(len(rooms)) - m.Offset
	}

	msg.Send200(a, t, &msg.GetRoomsRsp{RoomInfo: rooms[m.Offset : m.Offset+m.Limit]})
}
