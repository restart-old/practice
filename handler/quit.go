package handler

import "github.com/df-mc/dragonfly/server/entity/damage"

func (h PlayerHandler) HandleQuit() {
	if !h.p.CombatCD().Expired() {
		if src, ok := h.p.LastHurt(); ok {
			h.p.Kill(damage.SourceEntityAttack{Attacker: src})
		}
	}
	h.p.Server().RemovePlayer(h.p)
}
