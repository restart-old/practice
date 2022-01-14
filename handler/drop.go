package handler

import (
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
)

func (h *PlayerHandler) HandleItemDrop(ctx *event.Context, e *entity.Item) { ctx.Cancel() }
