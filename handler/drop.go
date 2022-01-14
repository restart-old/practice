package handler

import (
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
)

func (h *InventoryHandler) HandleDrop(ctx *event.Context, slot int, it item.Stack) {
	ctx.Cancel()
}

func (h *PlayerHandler) HandleItemDrop(ctx *event.Context, e *entity.Item) { ctx.Cancel() }
