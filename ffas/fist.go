package ffas

import (
	"Practice/custom"
	"Practice/kits"
	"fmt"

	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-plus/kit"
	"github.com/go-gl/mathgl/mgl64"
)

func FistFFA(s *custom.Server) Fist { return Fist{server: s} }

type Fist struct {
	server *custom.Server
}

func (Fist) Name() string { return "Fist" }
func (fist Fist) World() *world.World {
	worldName := fist.Name()
	w, ok := fist.server.WorldManager().World(worldName)
	if !ok {
		fmt.Printf("ffa: world '%s' not loaded\n", worldName)
	}
	return w.World
}
func (Fist) Position() mgl64.Vec3 { return mgl64.Vec3{220, 67, 210} }
func (fist Fist) Kit() kit.Kit {
	kitName := fist.Name()
	k, ok := kits.ByName(kitName)
	if !ok {
		fmt.Printf("ffa: kit '%s' not registered\n", kitName)
	}
	return k
}
