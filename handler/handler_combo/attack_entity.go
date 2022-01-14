package handler_combo

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
)

func (h handler_combo) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
	if !ctx.Cancelled() {
		h.p.AddCombo(1)
	}
}
