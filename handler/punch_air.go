package handler

import "github.com/df-mc/dragonfly/server/event"

func (h *PlayerHandler) HandlePunchAir(ctx *event.Context) {
	h.p.AddCPS(1)
	h.p.Player().SendTip(h.p.CPS())
}
