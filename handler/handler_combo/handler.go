package handler_combo

import (
	"github.com/RestartFU/practice/custom"
	"github.com/df-mc/dragonfly/server/player"
)

type handler_combo struct {
	player.NopHandler
	p *custom.Player
}

func (handler_combo) Name() string { return "handler_combo" }

func NewComboHandler(p *custom.Player) *handler_combo { return &handler_combo{p: p} }
