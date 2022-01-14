package handler

import (
	"time"

	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
)

func (h *PlayerHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
	p := h.p.Player()
	h.p.AddCPS(1)
	p.SendTip(h.p.CPS())
	if p.World() == h.p.Server().WorldManager().DefaultWorld() {
		ctx.Cancel()
		return
	}
	h.p.SetCombat(15 * time.Second)

	if h.p.CPS() >= 20 {
		p.Message("stop clicking so fast")
		ctx.Cancel()
	}
	*force = 0.398
	*height = 0.405
}
