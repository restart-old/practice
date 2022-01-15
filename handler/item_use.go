package handler

import (
	"math"
	"time"

	"github.com/RestartFU/practice/custom"

	"github.com/df-mc/dragonfly/server/event"
)

func (h *PlayerHandler) HandleItemUse(ctx *event.Context) {
	heldItem, _ := h.p.HeldItems()

	// Pearl
	pearl := custom.EnderPearItem{}
	if heldItem.Item() == pearl {
		if cd := h.p.PearlCD(); !cd.Expired() {
			ctx.Cancel()
			h.p.Messagef("§cYou're on pearl cooldown for %v seconds.", math.Round(cd.UntilExpiration().Seconds()))
			return
		} else {
			cd.SetCooldown(10 * time.Second)
		}
	}
}
