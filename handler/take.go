package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
)

func (h *InventoryHandler) HandleTake(ctx *event.Context, slot int, it item.Stack) {}
