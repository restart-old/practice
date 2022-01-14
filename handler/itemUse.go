package handler

import (
	"Practice/custom"
	"math"
	"time"

	"github.com/df-mc/dragonfly/server/event"
)

func (h *PlayerHandler) HandleItemUse(ctx *event.Context) {
	heldItem, _ := h.p.Player().HeldItems()

	// Pearl
	pearl := custom.EnderPearItem{}
	if heldItem.Item() == pearl {
		if cd, ok := h.p.PearlCD(); ok {
			ctx.Cancel()
			h.p.Player().Messagef("You're on pearl cooldown for %v seconds.", math.Round(time.Until(cd).Seconds()))
			return
		}
		h.p.SetPearlCD(10 * time.Second)
	}
}
