package custom

import (
	"math"
	"math/rand"
	"sync/atomic"

	"github.com/df-mc/dragonfly/server/entity/physics"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// Lightning is a lethal element to thunderstorms. Lightning momentarily increases the skylight's brightness to slightly greater than full daylight.
type Lightning struct {
	pos atomic.Value

	state    int
	liveTime int
}

// NewLightning creates a lightning entity. The lightning entity will be positioned at the position passed.
func NewLightning(pos mgl64.Vec3) *Lightning {
	li := &Lightning{
		state:    2,
		liveTime: rand.Intn(3) + 1,
	}
	li.pos.Store(mgl64.Vec3{math.Floor(pos[0]), math.Floor(pos[1]), math.Floor(pos[2])})

	return li
}

// Position returns the current position of the lightning entity.
func (li *Lightning) Position() mgl64.Vec3 {
	return li.pos.Load().(mgl64.Vec3)
}

// World returns the world that the lightning entity is currently in, or nil if it is not added to a world.
func (li *Lightning) World() *world.World {
	w, _ := world.OfEntity(li)
	return w
}

// AABB ...
func (Lightning) AABB() physics.AABB {
	return physics.NewAABB(mgl64.Vec3{}, mgl64.Vec3{})
}

// Close closes the lighting.
func (li *Lightning) Close() error {
	li.World().RemoveEntity(li)
	return nil
}

// OnGround ...
func (Lightning) OnGround() bool {
	return false
}

// Rotation ...
func (li *Lightning) Rotation() (yaw, pitch float64) {
	return 0, 0
}

// EncodeEntity ...
func (li *Lightning) EncodeEntity() string {
	return "minecraft:lightning_bolt"
}

// Name ...
func (li *Lightning) Name() string {
	return "Lightning Bolt"
}

// New strikes the Lightning at a specific position in a new world.
func (li *Lightning) New(pos mgl64.Vec3) world.Entity {
	return NewLightning(pos)
}

// Tick ...
func (li *Lightning) Tick(w *world.World, _ int64) {
	pos := li.Position()

	if li.state == 2 { // Init phase
		w.PlaySound(pos, sound.Thunder{})

	}

	if li.state--; li.state < 0 {
		if li.liveTime == 0 {
			_ = li.Close()
		} else if li.state < -rand.Intn(10) {
			li.liveTime--
			li.state = 1
		}
	}
}

// DecodeNBT does nothing.
func (li *Lightning) DecodeNBT(map[string]interface{}) interface{} {
	return nil
}

// EncodeNBT does nothing.
func (li *Lightning) EncodeNBT() map[string]interface{} {
	return map[string]interface{}{}
}

// fire returns a fire block.
func fire() world.Block {
	f, ok := world.BlockByName("minecraft:fire", map[string]interface{}{"age": int32(0)})
	if !ok {
		panic("could not find fire block")
	}
	return f
}
