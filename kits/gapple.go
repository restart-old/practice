package kits

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/armour"
	"github.com/df-mc/dragonfly/server/item/enchantment"
	"github.com/df-mc/dragonfly/server/item/tool"
	"github.com/df-plus/kit"
)

type Gapple struct{}

func (Gapple) Name() string { return "§8Gapple" }
func (Gapple) Items() kit.Items {
	return kit.Items{
		kit.Set{Slot: 0}: item.NewStack(item.Sword{Tier: tool.TierDiamond}, 1).WithEnchantments(enchantment.Unbreaking{}.WithLevel(10)),
		kit.Set{Slot: 1}: item.NewStack(item.GoldenApple{}, 16),
	}
}
func (Gapple) Armour() [4]item.Stack {
	unbreaking := enchantment.Unbreaking{}.WithLevel(10)
	return [4]item.Stack{
		item.NewStack(item.Helmet{Tier: armour.TierDiamond}, 1).WithEnchantments(unbreaking),
		item.NewStack(item.Chestplate{Tier: armour.TierDiamond}, 1).WithEnchantments(unbreaking),
		item.NewStack(item.Leggings{Tier: armour.TierDiamond}, 1).WithEnchantments(unbreaking),
		item.NewStack(item.Boots{Tier: armour.TierDiamond}, 1).WithEnchantments(unbreaking),
	}
}
