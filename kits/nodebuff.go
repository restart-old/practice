package kits

import (
	"time"

	"github.com/RestartFU/practice/custom"

	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/enchantment"
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/df-plus/kit"
)

type NoDebuff struct{}

func (NoDebuff) Name() string { return "§8NoDebuff" }
func (NoDebuff) Items() kit.Items {
	return kit.Items{
		kit.Set{Slot: 0}: item.NewStack(item.Sword{Tier: item.ToolTierDiamond}, 1).WithEnchantments(enchantment.Unbreaking{}.WithLevel(10)),
		kit.Set{Slot: 1}: item.NewStack(custom.EnderPearItem{}, 16),
		kit.Add{}:        item.NewStack(custom.SplashPotionItem{Type: potion.StrongHealing()}, 36),
	}
}
func (NoDebuff) Armour() [4]item.Stack {
	unbreaking := enchantment.Unbreaking{}.WithLevel(10)
	return [4]item.Stack{
		item.NewStack(item.Helmet{Tier: item.ArmourTierDiamond}, 1).WithEnchantments(unbreaking),
		item.NewStack(item.Chestplate{Tier: item.ArmourTierDiamond}, 1).WithEnchantments(unbreaking),
		item.NewStack(item.Leggings{Tier: item.ArmourTierDiamond}, 1).WithEnchantments(unbreaking),
		item.NewStack(item.Boots{Tier: item.ArmourTierDiamond}, 1).WithEnchantments(unbreaking),
	}
}

func (NoDebuff) Effects() []effect.Effect {
	return []effect.Effect{effect.New(effect.Speed{}, 1, 30*time.Hour)}
}
