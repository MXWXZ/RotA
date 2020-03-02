package internal

import (
	"rota/conf"
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
	util.Handler(Skeleton, &msg.ChangeTeam{}, checkUser(handleChangeTeam))
}

var nowID = 1
var rooms []msg.RoomInfo

func countTeam(s []msg.RoomMember) [conf.MaxRoomTeam + 1]int {
	var ret [conf.MaxRoomTeam + 1]int
	for _, v := range s {
		ret[v.Team] += 1
	}
	return ret
}

func findRoomMember(a gate.Agent, r *msg.RoomInfo) int {
	for i, v := range r.Members {
		if v.ID == agents[a].ID {
			return i
		}
	}
	return -1
}

func findAgentRoom(a gate.Agent) int {
	if agents[a].Room == 0 {
		return -1
	}

	for i, v := range rooms {
		if v.ID == agents[a].Room {
			return i
		}
	}
	return -1
}

func findAgentMember(a gate.Agent) (int, int) {
	rid := findAgentRoom(a)
	mid := -1
	if rid != -1 {
		mid = findRoomMember(a, &rooms[rid])
	}
	return rid, mid
}

func handleChangeTeam(args []interface{}) {
	var m *msg.ChangeTeam
	a := util.GetArgs(args, &m)

	if !util.RequireParam(m.Team) || m.Team > conf.MaxRoomTeam {
		log.Error("User [%v] invalid param in %v", agents[a].ID, util.GetFuncName())
		return
	}

	if agents[a].Room == 0 {
		log.Error("User [%v] change team error: NotInRoom", agents[a].ID)
		msg.SendError(a, "切换阵营失败，你不在任何一个房间中")
		return
	}

	if agents[a].Status == 1 {
		log.Error("User [%v] change team error: GameStarted", agents[a].ID)
		msg.SendError(a, "切换阵营失败，游戏已经开始")
		return
	}

	if rn, mn := findAgentMember(a); rn != -1 && mn != -1 {
		if rooms[rn].Members[mn].Team == m.Team {
			log.Error("User [%v] change team error: AlreadyInTeam", agents[a].ID)
			msg.SendError(a, "切换阵营失败，你已经在阵营中了")
			return
		}
		cnt := countTeam(rooms[rn].Members)
		if rooms[rn].Type == msg.Room_Solo {
			if m.Team > 2 {
				log.Error("User [%v] invalid param in %v", agents[a].ID, util.GetFuncName())
				return
			}
			if cnt[m.Team] < 1 {
				rooms[rn].Members[mn].Team = m.Team
				var tmp = msg.RoomInfoRsp(rooms[rn])
				broadCastRoom(rooms[rn].ID, &tmp)
				log.Release("User [%v] change team [%v]", agents[a].ID, m.Team)
			} else {
				log.Error("User [%v] change team error: Full", agents[a].ID)
				msg.SendError(a, "切换阵营失败，对方阵营已满")
			}
		}
	} else {
		log.Error("User [%v] change team error: MemberNotFound", agents[a].ID)
		msg.SendError(a, "切换阵营失败，你不在房间中")
	}
}

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
		log.Error("User [%v] get ready error: NotInRoom", agents[a].ID)
		msg.SendError(a, "准备失败，你不在任何一个房间中")
		return
	}

	if agents[a].Status == 1 { // started
		log.Error("User [%v] get ready error: GameStarted", agents[a].ID)
		msg.SendError(a, "准备失败，游戏已经开始")
		return
	}

	if rn, mn := findAgentMember(a); rn != -1 && mn != -1 {
		if rooms[rn].Master == agents[a].ID { // master cant ready
			log.Error("User [%v] get ready error: MasterReady", agents[a].ID)
			msg.SendError(a, "准备失败，你现在是房主")
			return
		}
		rooms[rn].Members[mn].Ready = m.Ready
		var tmp = msg.RoomInfoRsp(rooms[rn])
		broadCastRoom(rooms[rn].ID, &tmp)
		log.Release("User [%v] get ready [%v]", agents[a].ID, m.Ready)
	} else {
		log.Error("User [%v] get ready error: MemberNotFound", agents[a].ID)
		msg.SendError(a, "准备失败，你不在房间中")
	}
}

func handleJoinRoom(args []interface{}) {
	var m *msg.JoinRoom
	a := util.GetArgs(args, &m)

	if !util.RequireParam(m.ID) {
		log.Error("User [%v] invalid param in %v", agents[a].ID, util.GetFuncName())
		return
	}

	if agents[a].Room != 0 && agents[a].Room != m.ID { // master will join after create, Room will be set in create
		log.Error("User [%v] join room error: AlreadyInRoom", agents[a].ID)
		msg.Send(a, &msg.JoinRoomRsp{Code: 2})
		return
	}

	for i, v := range rooms {
		if v.ID == m.ID {
			if agents[a].Room == m.ID { // already in just send info
				msg.Send(a, &msg.JoinRoomRsp{Code: 0, Info: v})
				return
			}
			if v.Capacity-v.Size > 0 { // not full
				rooms[i].Size += 1
				cnt := countTeam(v.Members)
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
				log.Release("User [%v] join room [%v]", agents[a].ID, rooms[i].ID)
			} else {
				log.Error("User [%v] join room error: Full", agents[a].ID)
				msg.Send(a, &msg.JoinRoomRsp{Code: 1})
			}
			return
		}
	}

	log.Error("User [%v] join room error: RoomNotFound", agents[a].ID)
	msg.Send(a, &msg.JoinRoomRsp{Code: 3})
}

func handleNewRoom(args []interface{}) {
	var m *msg.NewRoom
	a := util.GetArgs(args, &m)

	if !util.RequireParam(m.Name, m.Type) {
		log.Error("User [%v] invalid param in %v", agents[a].ID, util.GetFuncName())
		return
	}

	if agents[a].Room != 0 { // already in room
		log.Error("User [%v] new room error: AlreadyInRoom", agents[a].ID)
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
	log.Release("User [%v] create room [%v]", agents[a].ID, nowID-1)
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
	if rn, mn := findAgentMember(a); rn != -1 && mn != -1 {
		agents[a].Room = 0
		if rooms[rn].Size == 1 { // only one, delete room
			rooms = append(rooms[:rn], rooms[rn+1:]...)
			broadCastRoom(0, &msg.DeleteRoomRsp{ID: rid})
		} else { // have other people
			rooms[rn].Members = append(rooms[rn].Members[:mn], rooms[rn].Members[mn+1:]...)
			rooms[rn].Size -= 1
			if rooms[rn].Master == uid { // new master
				rooms[rn].Master = rooms[rn].Members[0].ID
				rooms[rn].Members[0].Ready = 0 // master clear ready
			}
			var tmp = msg.RoomInfoRsp(rooms[rn])
			broadCastRoom(0, &tmp)
			broadCastRoom(rid, &tmp)
		}
	} else {
		agents[a].Room = 0 // may be member not found, so force room=0
		log.Error("User [%v] exit room error: MemberNotFound", agents[a].ID)
	}
	log.Release("User [%v] exit room [%v]", agents[a].ID, rid)
}
