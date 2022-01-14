package handler_combo

import (
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/event"
)

func (h handler_combo) HandleHurt(ctx *event.Context, d *float64, src damage.Source) {
	if !ctx.Cancelled() {
		h.p.ResetCombo()
	}
}
