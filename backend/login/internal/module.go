package internal

import (
	"rota/base"
	"rota/db"

	"github.com/name5566/leaf/module"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	db.InitRDB(new(db.User))
	db.InitRS()
}

func (m *Module) OnDestroy() {
	db.RSClient.FlushDB()
	db.CloseRDB()
	db.CloseRS()
}
