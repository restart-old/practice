package handler

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/entity/damage"
)

func (h PlayerHandler) HandleQuit() {
	if src, ok := h.p.LastHurt(); ok && !h.p.CombatCD().Expired() {
		h.p.Kill(damage.SourceEntityAttack{Attacker: src})
	}
	h.p.Server().RemovePlayer(h.p)
	fmt.Printf("%s[%v] has left the server\n", h.p.Name(), h.p.Addr().String())
}
