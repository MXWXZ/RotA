package internal

import (
	"reflect"
	"rota/db"
	"rota/msg"
	"rota/util"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/name5566/leaf/log"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.Login{}, handleLogin)
	handler(&msg.Signup{}, handleSignup)
}

func checkUser(n string, p string) (*db.User, error) {
	if p != "" {
		p = util.HashPwd(p)
	}
	var usr db.User
	err := db.RDBClient.Where(&db.User{UserName: n, UserPass: p}).First(&usr).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &usr, nil
}

func handleSignup(args []interface{}) {
	var m *msg.Signup
	a, t := util.GetArgs(args, &m)

	if !util.RequireParam(m.UserName, m.UserPass) {
		msg.Send400(a, t)
		return
	}

	ret, err := checkUser(m.UserName, "")
	if err != nil {
		msg.Send500(a, t, err)
		return
	} else if ret != nil {
		msg.Send200(a, t, &msg.SignupRsp{Status: 1})
		return
	}

	db.RDBClient.Create(&db.User{UserName: m.UserName, UserPass: util.HashPwd(m.UserPass)})
	msg.Send200(a, t, &msg.SignupRsp{Status: 0})
	log.Release("user sign up \"%s\"", m.UserName)
}

func handleLogin(args []interface{}) {
	var m *msg.Login
	a, t := util.GetArgs(args, &m)

	if !util.RequireParam(m.UserName, m.UserPass) {
		msg.Send400(a, t)
		return
	}

	ret, err := checkUser(m.UserName, m.UserPass)
	if err != nil {
		msg.Send500(a, t, err)
		return
	} else if ret == nil {
		msg.Send200(a, t, &msg.LoginRsp{Status: 1})
		return
	}

	token := util.GetToken()
	db.RSClient.HMSet(token, map[string]interface{}{
		"ID":   ret.ID,
		"Name": ret.UserName,
	})
	db.RSClient.Expire(token, 1*time.Hour)
	msg.Send200(a, t, &msg.LoginRsp{
		Status: 0,
		ID:     ret.ID,
		Token:  token,
	})
	log.Release("user login \"%s\"", m.UserName)
}
