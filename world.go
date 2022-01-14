package main

import (
	"Practice/custom"
	"Practice/handler"
	"log"

	"github.com/df-mc/dragonfly/server/world"
)

func loadWorld(path, name string, s *custom.Server) {
	settings := &world.Settings{
		Name: name,
	}
	if err := s.WorldManager().LoadWorld(path, settings, world.Overworld); err != nil {
		log.Fatalln(err)
	}
}

func setWorldSettings(s *custom.Server) {
	for _, w := range s.WorldManager().Worlds() {
		w.Handle(handler.NewWorldHandler(s.World()))
		w.SetTime(0)
		w.StopTime()
		w.StopRaining()
		w.StopWeatherCycle()
		w.StopThundering()
		w.SetDifficulty(world.DifficultyHard)
	}
}
