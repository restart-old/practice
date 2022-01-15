package custom

import (
	"math"
	"time"

	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/entity/healing"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/tool"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-plus/cooldown"
	"github.com/df-plus/ffa"
	"github.com/df-plus/kit"
)

type Player struct {
	*player.Player
	s *Server

	combo int
	cps   int

	cdManager *cooldown.Manager

	lastHurt world.Entity

	ffa ffa.FFA
}

// NewPlayer returns a new *Player.
func NewPlayer(p *player.Player, s *Server) *Player {
	cPlayer := &Player{Player: p, s: s, cdManager: cooldown.NewManager()}
	s.AddPlayer(cPlayer)
	return cPlayer
}

// CDManager ...
func (p *Player) CDManager() *cooldown.Manager { return p.cdManager }

// Server returns the *custom.Server of the player.
func (p *Player) Server() *Server { return p.s }

// AddCPS adds the int passed to the player cps,
// and removes that same amount a second after.
func (p *Player) AddCPS(n int) {
	p.cps += n
	time.AfterFunc(1*time.Second, func() {
		p.cps -= n
	})
}

// CPS returns the current amount of cps of the player.
func (p *Player) CPS() int { return p.cps }

// Operator returns a bool of if the player is an operator or not.
func (p *Player) Operator() bool { return p.Server().operators.Listed(p.Name()) }

// Whitelisted returns a bool of if the player is whitelisted or not.
func (p *Player) Whitelisted() bool { return p.Server().whitelist.Listed(p.Name()) }

// Banned returns a bool of if the player is banned or not.
func (p *Player) Banned() bool { return p.Server().ban.Listed(p.Name()) }

// SetFFA ...
func (p *Player) SetFFA(ffa ffa.FFA) { p.ffa = ffa }

// FFA ...
func (p *Player) FFA() (ffa.FFA, bool) { return p.ffa, p.ffa != nil }

// ResetCombo ...
func (p *Player) ResetCombo() { p.combo = 0 }

// AddCombo ...
func (p *Player) AddCombo(n int) { p.combo += n }

// Combo ...
func (p *Player) Combo() int { return p.combo }

// LastHurt ...
func (p *Player) LastHurt() (world.Entity, bool) { return p.lastHurt, p.lastHurt != nil }

// SetLastHurt
func (p *Player) SetLastHurt(e world.Entity) {
	p.lastHurt = e
}

// RemoveAllEffects ...
func (p *Player) RemoveAllEffects() {
	for _, e := range p.Effects() {
		p.RemoveEffect(e.Type())
	}
}

// WouldDie ...
func (p *Player) WouldDie(damage float64) bool {
	return p.Health()-damage <= 0
}

// AddToWorld ...
func (p *Player) AddToWorld(w *world.World) {
	if p == nil || p.Player == nil || w == nil {
		return
	}
	p.World().RemoveEntity(p.Player)
	w.AddEntity(p.Player)
	p.Teleport(w.Spawn().Vec3())
}

func (p *Player) CombatCD() *cooldown.Cooldown { return p.cdManager.Cooldown("combat_logger") }
func (p *Player) ChatCD() *cooldown.Cooldown   { return p.cdManager.Cooldown("chat") }
func (p *Player) PearlCD() *cooldown.Cooldown  { return p.cdManager.Cooldown("ender_pearl") }

// Kill ...
func (p *Player) Kill(src damage.Source) {
	if src == nil && p == nil || p.Player == nil {
		return
	}
	switch src := src.(type) {
	case damage.SourceEntityAttack:
		if src.Attacker == nil {
			return
		}
		player, ok := p.Server().PlayerByName(src.Attacker.Name())
		if !ok {
			return
		}
		if m, ok := MessageFFA(p, player.Player); ok {
			chat.Global.WriteString(m)
		}
		player.CombatCD().SetCooldown(0)
		player.ReKit()
		p.Spawn()
	default:
		p.Spawn()
	}
}

// ReKit ...
func (p *Player) ReKit() {
	if p == nil || p.Player == nil {
		return
	}
	if cd := p.CombatCD(); !cd.Expired() {
		p.Messagef("§cYou're still in combat for %v seconds", math.Round(cd.UntilExpiration().Seconds()))
		return
	}
	p.Heal(p.MaxHealth(), healing.SourceCustom{})

	p.Inventory().Clear()
	p.Armour().Clear()

	k, ok := p.FFA()
	if !ok {
		p.Messagef("§cYou're not in any FFA, teleported to spawn")
		p.Spawn()
		return
	}
	kit.GiveKit(p.Player, k.Kit())

}

// Spawn ...
func (p *Player) Spawn() {
	if p == nil || p.Player == nil || p.World() == nil {
		return
	}
	if cd := p.CombatCD(); !cd.Expired() {
		p.World().AddEntity(NewLightning(p.Position()))
		cd.SetCooldown(0)
	}

	p.Inventory().Clear()
	p.Armour().Clear()

	p.AddToWorld(p.Server().World())
	p.RemoveAllEffects()
	p.Heal(p.MaxHealth(), healing.SourceCustom{})
	p.SetFFA(nil)
	p.Inventory().SetItem(0, item.NewStack(item.Sword{Tier: tool.TierDiamond}, 1).WithCustomName("§eFFA - Unranked"))
}
