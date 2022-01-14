package handler

import (
	"Practice/custom"
	"fmt"

	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/entity/healing"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-plus/kit"
)

func (h *PlayerHandler) HandleDeath(src damage.Source) {
	h.p.SetFFA(nil)
	if s, ok := src.(damage.SourceEntityAttack); ok {
		p1, ok := h.p.Server().PlayerByName(s.Attacker.Name())
		if ok {
			p2 := p1.Player()
			p2.Heal(p2.MaxHealth(), healing.SourceCustom{})

			p2.Inventory().Clear()
			p2.Armour().Clear()

			if k, ok := p1.FFA(); ok {
				fmt.Println(k.Kit().Name(), k)
				kit.GiveKit(p2, k.Kit())
			}
		}
	}

	p := h.p.Player()
	for _, e := range p.World().Entities() {
		if _, ok := e.(*player.Player); !ok {
			e.World().RemoveEntity(e)
		}
	}
	_, c := h.p.Combat()
	if _, ok := src.(damage.SourceCustom); !ok || c {
		p.World().AddEntity(custom.NewLightning(p.Position()))
	}
	h.p.SetCombat(0)
	p.Respawn()
}
