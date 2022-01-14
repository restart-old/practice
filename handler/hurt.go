package handler

import (
	"time"

	"github.com/RestartFU/practice/ffas"

	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/event"
)

func (h *PlayerHandler) HandleHurt(ctx *event.Context, d *float64, src damage.Source) {
	ctx.After(func(cancelled bool) {
		if !cancelled {
			if ffa, ok := h.p.FFA(); ok {
				switch ffa.(type) {
				case ffas.Fist:
					h.p.Player().SetAttackImmunity(450 * time.Millisecond)
				}
			}
		}
	})
	if h.p.Player().World() == h.p.Server().World() {
		if h.p.Player().Position().Y() < 20 {
			h.p.Player().Teleport(h.p.Player().World().Spawn().Vec3())
		}
		ctx.Cancel()
	}
	switch src.(type) {
	case damage.SourceEntityAttack:
		h.p.SetCombat(15 * time.Second)
	case damage.SourceFall:
		ctx.Cancel()
	}
}
