package ffas

import (
	"Practice/custom"
	"Practice/kits"
	"fmt"

	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-plus/kit"
	"github.com/go-gl/mathgl/mgl64"
)

func GappleFFA(s *custom.Server) Gapple { return Gapple{server: s} }

type Gapple struct {
	server *custom.Server
}

func (Gapple) Name() string { return "Gapple" }
func (gapple Gapple) World() *world.World {
	worldName := gapple.Name()
	w, ok := gapple.server.WorldManager().World(worldName)
	if !ok {
		fmt.Printf("ffa: world '%s' not loaded\n", worldName)
	}
	return w.World
}
func (Gapple) Position() mgl64.Vec3 { return mgl64.Vec3{220, 67, 210} }
func (gapple Gapple) Kit() kit.Kit {
	kitName := gapple.Name()
	k, ok := kits.ByName(kitName)
	if !ok {
		fmt.Printf("ffa: kit '%s' not registered\n", kitName)
	}
	return k
}
