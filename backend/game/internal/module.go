package internal

import (
	"rota/base"
	"rota/db"

	"github.com/name5566/leaf/module"
)

var (
	Skeleton = base.NewSkeleton()
	ChanRPC  = Skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = Skeleton
	db.InitRS()
}

func (m *Module) OnDestroy() {
	db.CloseRS()
}
