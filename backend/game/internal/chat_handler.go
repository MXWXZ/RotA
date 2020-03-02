package internal

import (
	"rota/msg"
	"rota/util"

	"github.com/name5566/leaf/log"
)

func init() {
	util.Handler(Skeleton, &msg.Chat{}, checkUser(handleChat))
}

func handleChat(args []interface{}) {
	var m *msg.Chat
	a := util.GetArgs(args, &m)

	if !util.RequireParam(m.Channel, m.Msg) {
		log.Error("User [%v] invalid param in %v", agents[a].ID, util.GetFuncName())
		return
	}

	ms := msg.ChatRsp{
		Name:    agents[a].Name,
		Channel: m.Channel,
		Msg:     m.Msg,
	}
	if m.Channel == msg.Chat_Room {
		broadCastRoom(agents[a].Room, &ms)
	}
}
