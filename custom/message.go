package custom

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	deathInventory   = "§e%s§7[§e%v§7] was slain by §e%s§7[§e%v§7]"
	deathNoInventory = "§e%s§7 was slain by §e%s"
)

func ItemInInv(i world.Item, p *player.Player) int {
	var n int
	for _, item := range p.Inventory().Items() {
		e1, _ := item.Item().EncodeItem()
		e2, _ := i.EncodeItem()
		if e1 == e2 {
			n++
		}
	}
	return n
}

func MessageFFA(p *Player, p1 *player.Player) (string, bool) {
	if p == nil || p1 == nil {
		return "", false
	}
	ffa, ok := p.FFA()
	if !ok {
		return "", false
	}
	switch ffa.Name() {
	case "§8NoDebuff":
		return fmt.Sprintf(deathInventory, p.Name(), ItemInInv(item.SplashPotion{}, p.Player), p1.Name(), ItemInInv(item.SplashPotion{}, p1)), true
	default:
		return fmt.Sprintf(deathNoInventory, p.Name(), p1.Name()), true
	}
}
