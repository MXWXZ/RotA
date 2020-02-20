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
	a, t := util.GetArgs(args, &m)

	if !util.RequireParam(m.ID, m.Token) {
		msg.Send400(a, t)
		return
	}

	res, err := db.RSClient.HMGet(m.Token, "ID", "Name").Result()
	if err != nil || res[0] == nil || res[1] == nil || strconv.Itoa(int(m.ID)) != res[0].(string) {
		agents[a].Reset()
		log.Release("user id \"%v\" token invalid", m.ID)
		msg.Send200(a, t, &msg.CheckTokenRsp{Status: 1})
		return
	}
	db.RSClient.Expire(m.Token, 1*time.Hour) // reset expire time

	id, err := strconv.Atoi(res[0].(string))
	if err != nil {
		log.Error("%v", err)
	}

	agents[a].ID = id
	agents[a].Name = res[1].(string)
	msg.Send200(a, t, &msg.CheckTokenRsp{Status: 0})
	log.Release("user \"%v\" token valid", agents[a].Name)
}
