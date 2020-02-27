package internal

import (
	"rota/db"
	"rota/msg"
	"rota/util"
	"strconv"
	"time"

	"github.com/name5566/leaf/log"
)

func init() {
	util.Handler(Skeleton, &msg.CheckToken{}, handleCheckToken)
}

func handleCheckToken(args []interface{}) {
	var m *msg.CheckToken
	a := util.GetArgs(args, &m)

	if !util.RequireParam(m.ID, m.Token) {
		return
	}

	res, err := db.RSClient.HMGet(m.Token, "ID", "Name").Result()
	if err != nil || res[0] == nil || res[1] == nil || strconv.Itoa(int(m.ID)) != res[0].(string) {
		agents[a] = nil
		log.Error("User [%v] token invalid", m.ID)
		msg.Send(a, &msg.CheckTokenRsp{Code: 1})
		return
	}
	db.RSClient.Expire(m.Token, 1*time.Hour) // reset expire time

	id, err := strconv.Atoi(res[0].(string))
	if err != nil {
		log.Error("%v", err)
		msg.SendError(a, "Login server error")
	}

	if _, ok := users[id]; !ok { // not login before
		users[id] = &User{
			ID:     id,
			Name:   res[1].(string),
			Room:   0,
			Status: 0,
		}
	}
	if users[id].Token != m.Token && users[id].Token != "" { // former token exist
		db.RSClient.Del(users[id].Token)
		delete(agents, users[id].Agent)
	}
	users[id].Agent = a
	users[id].Token = m.Token
	agents[a] = users[id]
	msg.Send(a, &msg.CheckTokenRsp{
		Code:   0,
		ID:     agents[a].ID,
		Name:   agents[a].Name,
		Room:   agents[a].Room,
		Status: agents[a].Status,
	})
	log.Release("User [%v] token valid", m.ID)
}
