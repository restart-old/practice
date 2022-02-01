package main

import (
	"github.com/RestartFU/gophig"
	"github.com/RestartFU/practice/commands"
	"github.com/RestartFU/practice/custom"
	"github.com/RestartFU/practice/handler"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
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

	for {
		p, err := s.Accept()
		if err != nil {
			return
		}
		p.Handle(handler.PlayerHandler(p))
	}
}
