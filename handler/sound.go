package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

func (h *WorldHandler) HandleSound(ctx *event.Context, s world.Sound, pos mgl64.Vec3) {
	attackFalse := sound.Attack{Damage: false}
	attackTrue := sound.Attack{Damage: true}
	punchAir := sound.Attack{}

	switch s {
	case attackFalse, attackTrue, punchAir:
		ctx.Cancel()
	}
}
