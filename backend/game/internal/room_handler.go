package internal

import (
	"rota/msg"
	"rota/util"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	util.Handler(Skeleton, &msg.GetRooms{}, checkUser(handleGetRooms))
	util.Handler(Skeleton, &msg.NewRoom{}, checkUser(handleNewRoom))
	util.Handler(Skeleton, &msg.JoinRoom{}, checkUser(handleJoinRoom))
	util.Handler(Skeleton, &msg.ReadyRoom{}, checkUser(handleReadyRoom))
	util.Handler(Skeleton, &msg.ExitRoom{}, checkUser(handleExitRoom))
}

var nowID = 1
var rooms []msg.RoomInfo

func handleExitRoom(args []interface{}) {
	var m *msg.ExitRoom
	a := util.GetArgs(args, &m)

	exitRoom(a)
}

func handleReadyRoom(args []interface{}) {
	var m *msg.ReadyRoom
	a := util.GetArgs(args, &m)

	if m.Ready != 0 && m.Ready != 1 {
		log.Error("User [%v] invalid param in %v", agents[a].ID, util.GetFuncName())
		return
	}

	if agents[a].Room == 0 { // not in room
		log.Error("User [%v] invalid ready", agents[a].ID)
		msg.SendError(a, "准备失败，你不在任何一个房间中")
		return
	}

	if agents[a].Status == 1 { // started
		log.Error("User [%v] invalid ready", agents[a].ID)
		msg.SendError(a, "准备失败，游戏已经开始")
		return
	}

	for i, v := range rooms {
		if v.ID == agents[a].Room {
			if v.Master == agents[a].ID { // master cant ready
				log.Error("User [%v] master invalid ready", agents[a].ID)
				msg.SendError(a, "准备失败，你现在是房主")
				return
			}
			for j, k := range rooms[i].Members {
				if k.ID == agents[a].ID {
					rooms[i].Members[j].Ready = m.Ready
					var tmp = msg.RoomInfoRsp(rooms[i])
					broadCastRoom(v.ID, &tmp)
				}
			}
			break
		}
	}
	log.Release("User [%v] get ready %v", agents[a].ID, m.Ready)
}

func handleJoinRoom(args []interface{}) {
	var m *msg.JoinRoom
	a := util.GetArgs(args, &m)

	if !util.RequireParam(m.ID) {
		log.Error("User [%v] missing param in %v", agents[a].ID, util.GetFuncName())
		return
	}

	if agents[a].Room != 0 && agents[a].Room != m.ID { // master will join after create, Room will be set in create
		log.Error("User [%v] join room invalid", agents[a].ID)
		msg.Send(a, &msg.JoinRoomRsp{Code: 2})
		return
	}

	for i, v := range rooms {
		if v.ID == m.ID {
			if agents[a].Room == m.ID { // already in just send info
				msg.Send(a, &msg.JoinRoomRsp{Code: 0, Info: rooms[i]})
				return
			}
			if v.Capacity-v.Size > 0 { // not full
				rooms[i].Size += 1
				cnt := util.SumSlice(rooms[i].Members)
				member := msg.RoomMember{
					ID:   agents[a].ID,
					Name: agents[a].Name,
					Team: 0,
				}
				if rooms[i].Type == msg.Room_Solo {
					if cnt[1] == 1 {
						member.Team = 2
					} else {
						member.Team = 1
					}
				}
				rooms[i].Members = append(rooms[i].Members, member)
				var tmp = msg.RoomInfoRsp(rooms[i])
				broadCastRoom(v.ID, &tmp)
				broadCastRoom(0, &tmp)
				agents[a].Room = v.ID
				msg.Send(a, &msg.JoinRoomRsp{Code: 0, Info: rooms[i]})
				log.Release("User [%v] join room [%v]", agents[a].ID, v.ID)
			} else {
				msg.Send(a, &msg.JoinRoomRsp{Code: 1})
			}
			return
		}
	}
	msg.Send(a, &msg.JoinRoomRsp{Code: 3})
}

func handleNewRoom(args []interface{}) {
	var m *msg.NewRoom
	a := util.GetArgs(args, &m)

	if !util.RequireParam(m.Name, m.Type) {
		log.Error("User [%v] missing param in %v", agents[a].ID, util.GetFuncName())
		return
	}

	if agents[a].Room != 0 { // already in room
		log.Error("User [%v] new room invalid", agents[a].ID)
		msg.SendError(a, "创建失败，你已经在一个房间中了")
		return
	}

	var rc int
	switch msg.RoomType(m.Type) {
	case msg.Room_Solo:
		rc = 1 * 2
	}
	info := msg.RoomInfo{
		ID:       nowID,
		Name:     m.Name,
		Type:     msg.RoomType(m.Type),
		Size:     1,
		Capacity: rc,
		Master:   agents[a].ID,
		Status:   0,
		Members: []msg.RoomMember{{
			ID:   agents[a].ID,
			Name: agents[a].Name,
			Team: 1,
		}},
	}
	rooms = append(rooms, info)
	nowID += 1
	agents[a].Room = info.ID

	var tmp = msg.NewRoomRsp(info)
	broadCastRoom(0, &tmp)
	msg.Send(a, &tmp)
	log.Release("User [%v] create room", agents[a].ID)
}

func handleGetRooms(args []interface{}) {
	var m *msg.GetRooms
	a := util.GetArgs(args, &m)

	if m.Limit < 0 || m.Offset < 0 {
		return
	}
	if m.Limit == 0 {
		m.Limit = 20
	} else if m.Limit > 20 {
		m.Limit = 20
	}
	if m.Offset >= len(rooms) {
		msg.Send(a, &msg.GetRoomsRsp{
			Total:    len(rooms),
			RoomInfo: make([]msg.RoomInfo, 0, 0),
		})
		return
	}
	if m.Offset+m.Limit > len(rooms) {
		m.Limit = len(rooms) - m.Offset
	}

	msg.Send(a, &msg.GetRoomsRsp{
		Total:    len(rooms),
		RoomInfo: rooms[m.Offset : m.Offset+m.Limit],
	})
	log.Release("User [%v] get rooms", agents[a].ID)
}

func exitRoom(a gate.Agent) {
	if agents[a] == nil || agents[a].Status == 1 || agents[a].Room == 0 {
		return
	}

	rid := agents[a].Room
	uid := agents[a].ID
	agents[a].Room = 0
	for i, v := range rooms {
		if v.ID == rid {
			if v.Size == 1 { // only one, delete room
				rooms = append(rooms[:i], rooms[i+1:]...)
				broadCastRoom(0, &msg.DeleteRoomRsp{ID: rid})
			} else { // have other people
				for ri, rm := range rooms[i].Members {
					if rm.ID == uid {
						rooms[i].Members = append(rooms[i].Members[:ri], rooms[i].Members[ri+1:]...)
					}
				}
				rooms[i].Size -= 1
				if rooms[i].Master == uid { // new master
					rooms[i].Master = rooms[i].Members[0].ID
				}
				var tmp = msg.RoomInfoRsp(rooms[i])
				broadCastRoom(0, &tmp)
				broadCastRoom(rid, &tmp)
			}
			break
		}
	}
	log.Release("User [%v] exit room [%v]", agents[a].ID, rid)
}
