package handler

import (
	"github.com/RestartFU/practice/custom"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type PlayerHandler struct {
	player.NopHandler
	p *custom.Player
}

func (*PlayerHandler) Name() string { return "Player Handler" }

func NewPlayerHandler(p *custom.Player) *PlayerHandler { return &PlayerHandler{p: p} }

type WorldHandler struct {
	world.NopHandler
	world *world.World
}

func NewWorldHandler(w *world.World) *WorldHandler { return &WorldHandler{world: w} }
