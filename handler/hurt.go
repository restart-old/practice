package handler

import (
	"time"

	"github.com/RestartFU/practice/ffas"
	"github.com/df-plus/kit"

	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/entity/healing"
	"github.com/df-mc/dragonfly/server/event"
)

func (h *PlayerHandler) HandleHurt(ctx *event.Context, d *float64, src damage.Source) {
	p := h.p.Player()

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
		return
	}

	switch src := src.(type) {
	case damage.SourceEntityAttack:
		if h.p.WouldDie(p.FinalDamageFrom(*d, src)) {
			ctx.Cancel()
			h.p.Spawn()

			p1, ok := h.p.Server().PlayerByName(src.Attacker.Name())
			if ok {
				p1.SetCombat(0)
				p2 := p1.Player()
				p2.Heal(p2.MaxHealth(), healing.SourceCustom{})

				p2.Inventory().Clear()
				p2.Armour().Clear()

				if k, ok := p1.FFA(); ok {
					kit.GiveKit(p2, k.Kit())
				}
			}
		} else {
			h.p.SetCombat(15 * time.Second)
		}
	case damage.SourceFall:
		ctx.Cancel()
	default:
		if h.p.WouldDie(p.FinalDamageFrom(*d, src)) {
			ctx.Cancel()
			h.p.Spawn()
		}
	}
}
