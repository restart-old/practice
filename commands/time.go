package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

// Time returns the time command
func Time() cmd.Command {
	return cmd.New("time", "set the time of your current world", nil,
		TimeSetWord{},
		TimeSetInt{},
	)
}

// TimeSetWord is the command used to set the time with a word such as "day" or "night".
type TimeSetWord struct {
	Sub  set
	Time time
}

// time is the enum argument used to change the time of the world using a word.
type time string

// Type ...
func (time) Type() string { return "Time" }

// Options ...
func (time) Options(cmd.Source) []string {
	return []string{"day", "night", "midnight", "midday", "sunrise", "sunset"}
}

// Run ...
func (command TimeSetWord) Run(src cmd.Source, out *cmd.Output) {
	time := map[time]int{
		"day":      1000,
		"night":    13000,
		"midnight": 18000,
		"midday":   6000,
		"sunrise":  23000,
		"sunset":   12000,
	}
	src.World().SetTime(time[command.Time])
}

// TimeSetInt is the command used to change the time of the world with an integer.
type TimeSetInt struct {
	Sub  set
	Time int
}

// Run ...
func (command TimeSetInt) Run(src cmd.Source, out *cmd.Output) {
	src.World().SetTime(command.Time)
}

// set is the SubCommand used to set the time of the world
type set string

// SubName ...
func (set) SubName() string { return "set" }
