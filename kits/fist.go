package kits

import (
	"time"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-plus/kit"
)

type Fist struct{}

func (Fist) Name() string { return "§8Fist" }
func (Fist) Items() kit.Items {
	return kit.Items{
		kit.Set{Slot: 0}: item.NewStack(item.Beef{Cooked: true}, 16),
	}
}
func (Fist) Armour() [4]item.Stack {
	return [4]item.Stack{
		{},
		{},
		{},
		{},
	}
}

func (Fist) Effects() []effect.Effect {
	return []effect.Effect{effect.New(effect.Resistance{}, 2, 30*time.Hour)}
}
