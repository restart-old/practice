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
					h.p.SetAttackImmunity(450 * time.Millisecond)
				}
			}
		}
	})

	if attSrc, ok := src.(damage.SourceEntityAttack); ok {
		h.p.SetLastHurt(attSrc.Attacker)
	}

	switch src.(type) {
	case damage.SourceVoid:
		h.p.Kill(src)
		ctx.Cancel()
	case damage.SourceFall:
		ctx.Cancel()
	default:
		if h.p.WouldDie(h.p.FinalDamageFrom(*d, src)) {
			h.p.Kill(src)
			ctx.Cancel()
		} else {
			h.p.CombatCD().SetCooldown(15 * time.Second)
		}
	}
}
