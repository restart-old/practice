package commands

import (
	"github.com/RestartFU/practice/custom"

	"github.com/df-mc/dragonfly/server/cmd"
)

type BAN struct {
	Player string
	server *custom.Server
}

func BanRunnable(server *custom.Server) cmd.Runnable { return BAN{server: server} }

func (op BAN) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {

		if target, ok := op.server.PlayerByName(op.Player); ok {
			target.Disconnect("You are now banned")
		}
		p.Server().Ban().Add(op.Player)
		p.Messagef("%s is now banned", op.Player)
	}
}

func (o BAN) Allow(src cmd.Source) bool {
	return o.server.Operators().Listed(src.Name())
}

type UNBAN struct {
	Player string
	server *custom.Server
}

func UNBANRunnable(server *custom.Server) cmd.Runnable { return UNBAN{server: server} }

func (op UNBAN) Run(src cmd.Source, output *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {
		if p.Server().Ban().Listed(op.Player) {
			p.Server().Ban().Remove(op.Player)
			p.Messagef("%s is no longer banned", op.Player)
		} else {
			output.Errorf("player %s is not banned", op.Player)
		}
	}
}

func (d UNBAN) Allow(src cmd.Source) bool {
	return d.server.Operators().Listed(src.Name())
}
