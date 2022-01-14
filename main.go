package main

import (
	"Practice/custom"
	"Practice/ffas"
	"Practice/handler"
	"Practice/kits"

	"github.com/RestartFU/gophig"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/tool"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-plus/items"
	"github.com/df-plus/moreHandlers"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sirupsen/logrus"
)

func init() {
	chat.Global.Subscribe(chat.StdoutSubscriber{})
}

func init() {
	kits.Register(kits.NoDebuff{})
	kits.Register(kits.Gapple{})
	kits.Register(kits.Fist{})
}

func main() {
	config := server.DefaultConfig()
	gophig.GetConfComplex("./config.toml", gophig.TOMLMarshaler{}, &config)

	s := custom.NewServer(&config, logger())

	ffas.Register(ffas.NoDebuffFFA(s))
	ffas.Register(ffas.GappleFFA(s))
	ffas.Register(ffas.FistFFA(s))

	RegisterAllCommands(s)

	s.World().SetSpawn(cube.PosFromVec3(mgl64.Vec3{291, 77, 176}))

	items.Register(FFASword{server: s})

	s.Allow(Allower{s: s})
	s.Start()
	s.CloseOnProgramEnd()

	loadWorld("./data/world2", "NoDebuff", s)
	loadWorld("./data/world3", "Gapple", s)
	loadWorld("./data/world4", "Fist", s)

	setWorldSettings(s)
	for {
		p, err := s.Accept()
		if err != nil {
			return
		}
		p.Player().Inventory().Clear()
		p.Player().Armour().Clear()
		go handleJoin(p)
	}

}

func logger() *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel
	return log
}
func handleJoin(p *custom.Player) {
	p.Player().SetGameMode(gamemode{})

	p.Player().Inventory().Handle(handler.NewInventoryHandler(p))

	h := moreHandlers.NewPlayerHandler()
	p.Player().Handle(h)

	h.AddHandler(handler.NewPlayerHandler(p))
	h.AddHandler(items.NewPlayerHandler(p.Player()))

	p.Player().Inventory().SetItem(0, item.NewStack(item.Sword{Tier: tool.TierDiamond}, 1).WithCustomName("§eFFA - Unranked"))
}
