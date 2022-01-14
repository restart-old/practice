package commands

import (
	"Practice/custom"

	"github.com/df-mc/dragonfly/server/cmd"
)

type OP struct {
	Player string
	server *custom.Server
}

func OPRunnable(server *custom.Server) cmd.Runnable { return OP{server: server} }

func (op OP) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {
		p.Server().Operators().Add(op.Player)
		p.Player().Messagef("%s is now OP", op.Player)
	}
}

func (o OP) Allow(src cmd.Source) bool {
	return o.server.Operators().Listed(src.Name())
}

type DEOP struct {
	Player string
	server *custom.Server
}

func DEOPRunnable(server *custom.Server) cmd.Runnable { return DEOP{server: server} }

func (op DEOP) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {
		p.Server().Operators().Remove(op.Player)
		p.Player().Messagef("%s is no longer OP", op.Player)
	}
}

func (d DEOP) Allow(src cmd.Source) bool {
	return d.server.Operators().Listed(src.Name())
}
