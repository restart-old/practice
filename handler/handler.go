package handler

import (
	"github.com/RestartFU/practice/custom"

	"github.com/df-mc/dragonfly/server/item/inventory"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type PlayerHandler struct {
	player.NopHandler
	p *custom.Player
}

func (*PlayerHandler) Name() string { return "Player Handler" }

func NewPlayerHandler(p *custom.Player) *PlayerHandler { return &PlayerHandler{p: p} }

type InventoryHandler struct {
	inventory.NopHandler
	inv *inventory.Inventory
}

func NewInventoryHandler(p *custom.Player) *InventoryHandler {
	return &InventoryHandler{inv: p.Player().Inventory()}
}

type WorldHandler struct {
	world.NopHandler
	world *world.World
}

func NewWorldHandler(w *world.World) *WorldHandler { return &WorldHandler{world: w} }
