package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

func (h *WorldHandler) HandleSound(ctx *event.Context, s world.Sound, pos mgl64.Vec3) {
	switch s.(type) {
	case sound.Attack:
		ctx.Cancel()
	}
}
