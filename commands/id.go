package commands

import (
	"github.com/RestartFU/practice/custom"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type ID struct {
	server *custom.Server
}

func IDRunnable(server *custom.Server) cmd.Runnable { return ID{server: server} }

func (ID) Run(src cmd.Source, o *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {
		heldItem, _ := p.Player().HeldItems()
		if heldItem.Empty() {
			o.Error("The item you're currently holding is not registered")
			return
		}
		rid, meta, ok := world.ItemRuntimeID(heldItem.Item())

		if !ok {
			o.Error("The item you're currently holding is not registered")
			return
		}
		p.Player().Messagef("%v:%v", rid, meta)
	}
}
func (i ID) Allow(src cmd.Source) bool {
	return i.server.Operators().Listed(src.Name())
}
