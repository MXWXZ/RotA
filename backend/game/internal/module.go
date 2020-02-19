package internal

import (
	"rota/base"

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
}

func (m *Module) OnDestroy() {

}
