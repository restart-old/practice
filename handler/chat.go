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
	if t, ok := h.p.ChatCD(); ok {
		h.p.Player().Messagef("You're on chat cooldown for %v seconds", math.Round(time.Until(t).Seconds()))
		return
	}
	if !h.p.Operator() {
		h.p.SetChatCD(3 * time.Second)
	}
	chat.Global.WriteString(fmt.Sprintf("§e%s: §f%s", h.p.Name(), *message))
}
