package custom

import (
	"image/color"
	"time"

	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/entity/physics"
	"github.com/df-mc/dragonfly/server/entity/physics/trace"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/particle"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// SplashPotion is an item that grants effects when thrown.
type SplashPotion struct {
	transform
	yaw, pitch float64

	age   int
	close bool

	owner world.Entity

	t potion.Potion
	c *entity.ProjectileComputer
}

// NewSplashPotion ...
func NewSplashPotion(pos mgl64.Vec3, yaw, pitch float64, owner world.Entity, t potion.Potion) *SplashPotion {
	s := &SplashPotion{
		yaw:   yaw,
		pitch: pitch,
		owner: owner,

		t: t,
		c: &entity.ProjectileComputer{MovementComputer: &entity.MovementComputer{
			Gravity:           0.05,
			Drag:              0.01,
			DragBeforeGravity: true,
		}},
	}
	s.transform = newTransform(s, pos)

	return s
}

// Name ...
func (s *SplashPotion) Name() string {
	return "Splash Potion"
}

// EncodeEntity ...
func (s *SplashPotion) EncodeEntity() string {
	return "minecraft:splash_potion"
}

// AABB ...
func (s *SplashPotion) AABB() physics.AABB {
	return physics.NewAABB(mgl64.Vec3{-0.125, 0, -0.125}, mgl64.Vec3{0.125, 0.25, 0.125})
}

// Rotation ...
func (s *SplashPotion) Rotation() (float64, float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.yaw, s.pitch
}

// Type returns the type of potion the splash potion will grant effects for when thrown.
func (s *SplashPotion) Type() potion.Potion {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.t
}

// Tick ...
func (s *SplashPotion) Tick(w *world.World, current int64) {
	if s.close {
		_ = s.Close()
		return
	}
	s.mu.Lock()
	m, result := s.c.TickMovement(s, s.pos, s.vel, s.yaw, s.pitch, s.ignores)
	yaw, pitch := m.Rotation()
	s.pos, s.vel, s.yaw, s.pitch = m.Position(), m.Velocity(), yaw, pitch
	s.mu.Unlock()

	s.age++
	m.Send()

	if m.Position()[1] < float64(w.Range()[0]) && current%10 == 0 {
		s.close = true
		return
	}

	if result != nil {
		aabb := s.AABB().Translate(m.Position())

		colour := color.RGBA{R: 0x38, G: 0x5d, B: 0xc6, A: 0xff}
		if effects := s.t.Effects(); len(effects) > 0 {
			colour, _ = effect.ResultingColour(effects)

			ignore := func(e world.Entity) bool {
				_, living := e.(entity.Living)
				return !living || e == s
			}

			for _, e := range w.EntitiesWithin(aabb.GrowVec3(mgl64.Vec3{8.25, 15, 8.25}), ignore) {
				pos := e.Position()
				if !e.AABB().Translate(pos).IntersectsWith(aabb.GrowVec3(mgl64.Vec3{4.125, 2.125, 4.125})) {
					continue
				}
				potency := float64(1)
				dist := world.Distance(pos, m.Position())
				if dist > 4 {
					continue
				}
				potency -= dist / 10
				splashed := e.(entity.Living)

				for _, eff := range effects {
					if p, ok := eff.Type().(effect.PotentType); ok {
						splashed.AddEffect(effect.NewInstant(p.WithPotency(potency), eff.Level()))
						continue
					}

					dur := time.Duration(float64(eff.Duration()) * 0.75 * potency)
					if dur < time.Second {
						continue
					}
					splashed.AddEffect(effect.New(eff.Type().(effect.LastingType), eff.Level(), dur))
				}
			}
		} else if s.t.Equals(potion.Water()) {
			switch result := result.(type) {
			case trace.BlockResult:
				pos := result.BlockPosition().Side(result.Face())
				if w.Block(pos) == fire() {
					w.SetBlock(pos, air())
				}

				for _, f := range cube.HorizontalFaces() {
					if h := pos.Side(f); w.Block(h) == fire() {
						w.SetBlock(h, air())
					}
				}
			case trace.EntityResult:
				// TODO: Damage endermen, blazes, striders and snow golems when implemented and rehydrate axolotls.
			}
		}

		w.AddParticle(m.Position(), particle.Splash{Colour: colour})
		w.PlaySound(m.Position(), sound.GlassBreak{})

		s.close = true
	}
}

// ignores returns whether the SplashPotion should ignore collision with the entity passed.
func (s *SplashPotion) ignores(e world.Entity) bool {
	_, ok := e.(entity.Living)
	return !ok || e == s || (s.age < 5 && e == s.owner)
}

// New creates a SplashPotion with the position, velocity, yaw, and pitch provided. It doesn't spawn the SplashPotion,
// only returns it.
func (s *SplashPotion) New(pos, vel mgl64.Vec3, yaw, pitch float64, t potion.Potion) world.Entity {
	splash := NewSplashPotion(pos, yaw, pitch, nil, t)
	splash.vel = vel
	return splash
}

// Owner ...
func (s *SplashPotion) Owner() world.Entity {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.owner
}

// Own ...
func (s *SplashPotion) Own(owner world.Entity) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.owner = owner
}

// DecodeNBT decodes the properties in a map to a SplashPotion and returns a new SplashPotion entity.
func (s *SplashPotion) DecodeNBT(data map[string]interface{}) interface{} {
	return nil
}

// EncodeNBT encodes the SplashPotion entity's properties as a map and returns it.
func (s *SplashPotion) EncodeNBT() map[string]interface{} {
	return map[string]interface{}{}
}

// air returns an air block.
func air() world.Block {
	f, ok := world.BlockByName("minecraft:air", map[string]interface{}{})
	if !ok {
		panic("could not find air block")
	}
	return f
}

// SplashPotion is an item that grants effects when thrown.
type SplashPotionItem struct {
	// Type is the type of splash potion.
	Type potion.Potion
}

// MaxCount ...
func (s SplashPotionItem) MaxCount() int {
	return 1
}

// Use ...
func (s SplashPotionItem) Use(w *world.World, user item.User, ctx *item.UseContext) bool {

	splash := SplashPotion{}

	yaw, pitch := user.Rotation()
	e := splash.New(entity.EyePosition(user), entity.DirectionVector(user).Mul(0.5), yaw, pitch, s.Type)
	if o, ok := e.(owned); ok {
		o.Own(user)
	}

	ctx.SubtractFromCount(1)

	w.PlaySound(user.Position(), sound.ItemThrow{})

	w.AddEntity(e)

	return true
}

// EncodeItem ...
func (s SplashPotionItem) EncodeItem() (name string, meta int16) {
	return "minecraft:splash_potion", int16(s.Type.Uint8())
}

type owned interface {
	world.Entity
	Owner() world.Entity
	Own(owner world.Entity)
}
