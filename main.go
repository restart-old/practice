package main

import (
	"github.com/RestartFU/gophig"
	"github.com/RestartFU/practice/commands"
	"github.com/RestartFU/practice/custom"
	"github.com/RestartFU/practice/handler"
	"github.com/df-mc/dragonfly/server"
<<<<<<< Updated upstream
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/tool"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-plus/items"
	"github.com/df-plus/moreHandlers"
	"github.com/go-gl/mathgl/mgl64"
=======
	"github.com/df-mc/dragonfly/server/cmd"
>>>>>>> Stashed changes
	"github.com/sirupsen/logrus"
)

func init() {
}

func main() {
	var config server.Config
	if err := gophig.GetConfComplex("./config.toml", gophig.TOMLMarshaler{}, &config); err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	logger.Level = logrus.DebugLevel

	s := custom.NewServer(&config, logger)

	cmd.Register(commands.WhiteList(s.Whitelist()))
	cmd.Register(commands.Time())
	cmd.Register(commands.GameMode())

	if err := s.Start(); err != nil {
		panic(err)
	}

	s.CloseOnProgramEnd()

<<<<<<< Updated upstream
	loadWorld("./data/world2", "NoDebuff", s)
	loadWorld("./data/world3", "Gapple", s)
	loadWorld("./data/world4", "Fist", s)

	setWorldSettings(s)
=======
>>>>>>> Stashed changes
	for {
		p, err := s.Accept()
		if err != nil {
			return
		}
<<<<<<< Updated upstream
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

	h := moreHandlers.NewPlayerHandler()
	p.Player().Handle(h)

	h.AddHandler(handler.NewPlayerHandler(p))
	h.AddHandler(items.NewPlayerHandler(p.Player()))
	h.AddHandler(handler_combo.NewComboHandler(p))

	p.Player().Inventory().SetItem(0, item.NewStack(item.Sword{Tier: tool.TierDiamond}, 1).WithCustomName("§eFFA - Unranked"))
=======
		p.Handle(handler.PlayerHandler(p))
	}
>>>>>>> Stashed changes
}
