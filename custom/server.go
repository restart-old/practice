package custom

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/RestartFU/gophig"
	"github.com/RestartFU/list"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	mw "github.com/df-plus/worldmanager"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*server.Server
	worldManager *mw.WorldManager

	whitelist *list.List
	operators *list.List
	ban       *list.List

	players   map[*player.Player]*Player
	playersMu sync.RWMutex
}

// newList ...
func newList(name, extension string) *list.List {
	settings := &list.Settings{
		CacheOnly: false,
		Gophig:    gophig.NewGophig(name, extension, 0777),
	}
	list, err := list.New(settings)
	if err != nil {
		fmt.Println(err)
	}
	return list
}

// NewServer returns a new *Server.
func NewServer(c *server.Config, log *logrus.Logger) *Server {
	s := server.New(c, log)
	worldManager := mw.New(s, log)

	ban := newList("./data/lists/bans", "toml")
	whitelist := newList("./data/lists/whitelist", "toml")
	operators := newList("./data/lists/operators", "toml")

	return &Server{Server: s,
		worldManager: worldManager,
		whitelist:    whitelist,
		operators:    operators,
		ban:          ban,
		players:      make(map[*player.Player]*Player),
	}
}

// WorldManager returns the world manager of the server.
func (s *Server) WorldManager() *mw.WorldManager { return s.worldManager }

// Accept accepts the incoming player and returns it as *custom.Player.
// it only returns an error when the server is closed.
func (s *Server) Accept() (*Player, error) {
	p, err := s.Server.Accept()
	return NewPlayer(p, s), err
}

// PlayerByName ...
func (s *Server) PlayerByName(username string) (*Player, bool) {
	for _, p := range s.Players() {
		if strings.EqualFold(p.Name(), username) {
			s.playersMu.RLock()
			defer s.playersMu.RUnlock()
			player, ok := s.players[p]
			return player, ok
		}
	}
	return nil, false
}

// RemovePlayer ...
func (s *Server) RemovePlayer(p *Player) {
	s.playersMu.Lock()
	defer s.playersMu.Unlock()

	delete(s.players, p.Player)
}

// AddPlayer ...
func (s *Server) AddPlayer(p *Player) {
	s.playersMu.Lock()
	defer s.playersMu.Unlock()

	s.players[p.Player] = p
}

// Whitelist returns the whitelist of the server.
func (s *Server) Whitelist() *list.List { return s.whitelist }

// Operators returns the operators of the server.
func (s *Server) Operators() *list.List { return s.operators }

// Ban returns the banned list of the server.
func (s *Server) Ban() *list.List { return s.ban }

// CloseOnProgramEnd will close the server and its worlds on program end.
func (s *Server) CloseOnProgramEnd() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		if err := s.WorldManager().Close(); err != nil {
			fmt.Fprintf(os.Stdout, "error shutting down server: %v", err)
		}
		if err := s.Close(); err != nil {
			fmt.Fprintf(os.Stdout, "error shutting down server: %v", err)
		}
	}()
}
