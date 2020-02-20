package internal

import (
	"rota/msg"
	"rota/util"
)

func init() {
	util.Handler(Skeleton, &msg.GetRooms{}, checkUser(handleGetRooms))
	util.Handler(Skeleton, &msg.NewRoom{}, checkUser(handleNewRoom))
}

var rooms []msg.RoomInfo

func handleNewRoom(args []interface{}) {

}

func handleGetRooms(args []interface{}) {
	var m *msg.GetRooms
	a, t := util.GetArgs(args, &m)

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
