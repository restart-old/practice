package handler

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player/chat"
)

func (h PlayerHandler) HandleChat(ctx *event.Context, message *string) {
	ctx.Cancel()
	if !h.p.Operator() && strings.Contains("discord.gg/", *message) {
		return
	}
	if cd, ok := h.p.ChatCD(); ok && !cd.Expired() {
		h.p.Player().Messagef("You're on chat cooldown for %v seconds", math.Round(cd.UntilExpiration().Seconds()))
		return
	} else {
		if !h.p.Operator() {
			cd.SetCooldown(3 * time.Second)
		}
	}
	chat.Global.WriteString(fmt.Sprintf("§e%s: §f%s", h.p.Name(), *message))
}
