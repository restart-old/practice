package custom

import (
	"math"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/entity/healing"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/tool"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-plus/ffa"
	"github.com/df-plus/kit"
	"github.com/go-gl/mathgl/mgl64"
)

type Player struct {
	player *player.Player
	s      *Server

	combo int
	cps   int

	pearlCD time.Time
	chatCD  time.Time
	combat  time.Time

	lastHurt world.Entity

	ffa ffa.FFA
}

// NewPlayer returns a new *Player.
func NewPlayer(p *player.Player, s *Server) *Player {
	cPlayer := &Player{player: p, s: s}
	s.AddPlayer(cPlayer)
	return cPlayer
}

// Name ...
func (p *Player) Name() string { return p.player.Name() }

// Position ...
func (p *Player) Position() mgl64.Vec3 { return p.player.Position() }

// SendCommandOutput ...
func (p *Player) SendCommandOutput(output *cmd.Output) {
	p.player.SendCommandOutput(output)
}

// World ...
func (p *Player) World() *world.World { return p.player.World() }

// Server returns the *custom.Server of the player.
func (p *Player) Server() *Server { return p.s }

// Combat returns the last time the user was in combat,
// and if the player is still in combat cooldown.
func (p *Player) Combat() (time.Time, bool) {
	return p.combat, p.combat.After(time.Now())
}

// Combat sets the combat for the player with the duration passed.
func (p *Player) SetCombat(d time.Duration) { p.combat = time.Now().Add(d) }

// PearlCD returns the last time the user threw a pearl,
// and if the player is still on pearl cooldown.
func (p *Player) PearlCD() (time.Time, bool) {
	return p.pearlCD, p.pearlCD.After(time.Now())
}

// SetPearlCD sets the pearl cooldown for the player with the duration passed.
func (p *Player) SetPearlCD(d time.Duration) { p.pearlCD = time.Now().Add(d) }

// SetChatCD sets the chat cooldown for the player with the duration passed.
func (p *Player) SetChatCD(d time.Duration) { p.chatCD = time.Now().Add(d) }

// ChatCD returns the last time the user sent a message,
// and if the player is still on chat cooldown.
func (p *Player) ChatCD() (time.Time, bool) {
	return p.chatCD, p.chatCD.After(time.Now())
}

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

// Player returns the *player.Player of the *Player.
func (p *Player) Player() *player.Player { return p.player }

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
	for _, e := range p.Player().Effects() {
		p.Player().RemoveEffect(e.Type())
	}
}

// WouldDie ...
func (p *Player) WouldDie(damage float64) bool {
	return p.Player().Health()-damage <= 0
}

// AddToWorld ...
func (p *Player) AddToWorld(w *world.World) {
	p.player.World().RemoveEntity(p.player)
	time.AfterFunc(1, func() {
		w.AddEntity(p.player)
		p.player.Teleport(w.Spawn().Vec3())
	})
}

// Kill ...
func (p *Player) Kill(src damage.Source) {
	if src == nil {
		return
	}
	if p == nil {
		return
	}
	switch src := src.(type) {

	case damage.SourceEntityAttack:
		player, ok := p.Server().PlayerByName(src.Attacker.Name())
		if !ok {
			return
		}
		if m, ok := MessageFFA(p, player.player); ok {
			chat.Global.WriteString(m)
		}
		player.SetCombat(0)
		player.ReKit()
		p.Spawn()
	default:
		p.Spawn()
	}
}

// ReKit ...
func (p *Player) ReKit() {
	if p == nil {
		return
	}
	player := p.Player()
	if t, ok := p.Combat(); ok {
		player.Messagef("§cYou're still in combat for %v seconds", math.Round(time.Until(t).Seconds()))
		return
	}
	player.Heal(player.MaxHealth(), healing.SourceCustom{})

	player.Inventory().Clear()
	player.Armour().Clear()

	k, ok := p.FFA()
	if !ok {
		player.Messagef("§cYou're not in any FFA, teleported to spawn")
		p.Spawn()
		return
	}
	kit.GiveKit(player, k.Kit())

}

// Spawn ...
func (p *Player) Spawn() {
	if p == nil {
		return
	}
	player := p.player

	for _, e := range p.World().Entities() {
		if _, ok := e.(world.Item); ok {
			e.World().RemoveEntity(e)
		}
	}
	if _, c := p.Combat(); c {
		p.World().AddEntity(NewLightning(p.Position()))
	}
	p.SetCombat(0)

	player.Inventory().Clear()
	player.Armour().Clear()

	p.AddToWorld(p.Server().World())
	p.RemoveAllEffects()
	p.Player().Heal(player.MaxHealth(), healing.SourceCustom{})
	p.SetFFA(nil)
	p.Player().Inventory().SetItem(0, item.NewStack(item.Sword{Tier: tool.TierDiamond}, 1).WithCustomName("§eFFA - Unranked"))
}
