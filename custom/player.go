package custom

import (
<<<<<<< Updated upstream
	"math"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/damage"
	"github.com/df-mc/dragonfly/server/entity/healing"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/tool"
=======
	cooldown2 "github.com/RestartFU/practice/cooldown"
>>>>>>> Stashed changes
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-plus/cooldown"
<<<<<<< Updated upstream
	"github.com/df-plus/ffa"
	"github.com/df-plus/kit"
	"github.com/go-gl/mathgl/mgl64"
=======
>>>>>>> Stashed changes
)

// Player embeds *player.Player, so we can have more information into the struct.
type Player struct {
<<<<<<< Updated upstream
	player *player.Player
	s      *Server

	combo int
	cps   int

	cdManager *cooldown.Manager
=======
	// This is the embedded *player.Player.
	*player.Player
>>>>>>> Stashed changes

	// The display name is the string that will be shown as the name of the player
	// in chat and in the name tag.
	displayName string

	// The coolDown manager manages the coolDowns for the Player
	coolDownManager *cooldown.Manager
}

// NewPlayer returns a new *Player.
<<<<<<< Updated upstream
func NewPlayer(p *player.Player, s *Server) *Player {
	cPlayer := &Player{player: p, s: s, cdManager: cooldown.NewManager()}
	cPlayer.cdManager.NewCooldown("ender_pearl")
	cPlayer.cdManager.NewCooldown("combat_logger")
	cPlayer.cdManager.NewCooldown("chat")
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

func (p *Player) CombatCD() (*cooldown.Cooldown, bool) { return p.cdManager.Cooldown("combat_logger") }
func (p *Player) ChatCD() (*cooldown.Cooldown, bool)   { return p.cdManager.Cooldown("chat") }
func (p *Player) PearlCD() (*cooldown.Cooldown, bool)  { return p.cdManager.Cooldown("ender_pearl") }

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
		if combat, ok := player.CombatCD(); ok {
			combat.SetCooldown(0)
		}
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
	if cd, ok := p.CombatCD(); ok && !cd.Expired() {
		player.Messagef("§cYou're still in combat for %v seconds", math.Round(cd.UntilExpiration().Seconds()))
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
	if cd, c := p.CombatCD(); c && !cd.Expired() {
		p.World().AddEntity(NewLightning(p.Position()))
		cd.SetCooldown(0)
	}

	player.Inventory().Clear()
	player.Armour().Clear()

	p.AddToWorld(p.Server().World())
	p.RemoveAllEffects()
	p.Player().Heal(player.MaxHealth(), healing.SourceCustom{})
	p.SetFFA(nil)
	p.Player().Inventory().SetItem(0, item.NewStack(item.Sword{Tier: tool.TierDiamond}, 1).WithCustomName("§eFFA - Unranked"))
=======
func NewPlayer(p *player.Player) *Player {
	return &Player{
		Player:          p,
		displayName:     p.Name(),
		coolDownManager: cooldown.NewManager(),
	}
}

// DisplayName returns the display name of the player.
func (player *Player) DisplayName() (displayName string) { return player.displayName }

// SetDisplayName sets a new display name for the player.
// This display name will be shown in chat and in the name tag.
func (player *Player) SetDisplayName(displayName string) { player.displayName = displayName }

// CoolDown returns the *cooldown.Cooldown of the argument passed.
func (player *Player) CoolDown(cd cooldown2.CoolDown) *cooldown.Cooldown {
	return player.coolDownManager.Cooldown(cd.String())
>>>>>>> Stashed changes
}
