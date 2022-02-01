package custom

import (
	"github.com/RestartFU/gophig"
	"github.com/RestartFU/list"
	"github.com/df-mc/dragonfly/server"
	"github.com/sirupsen/logrus"
)

// Server embeds *server.Server.
type Server struct {
	*server.Server
	logger *logrus.Logger

	whitelist *list.List
}

// NewServer returns a new *Server
func NewServer(config *server.Config, logger *logrus.Logger) *Server {
	s := server.New(config, logger)
	wl, err := list.New(&list.Settings{CacheOnly: false, Gophig: gophig.NewGophig("./whitelist", "toml", 0777)})
	if err != nil {
		panic(err)
	}
	return &Server{
		Server:    s,
		logger:    logger,
		whitelist: wl,
	}
}

// Accept accepts the incoming connections and returns a new *custom.Player.
// It only returns an error when the server is closed.
func (server *Server) Accept() (*Player, error) {
	p, err := server.Server.Accept()
	if err != nil {
		return nil, err
	}
	return NewPlayer(p), nil
}

<<<<<<< Updated upstream
// RemovePlayer ...
func (s *Server) RemovePlayer(p *Player) {
	s.playersMu.Lock()
	defer s.playersMu.Unlock()

	delete(s.players, p.player)
}

// AddPlayer ...
func (s *Server) AddPlayer(p *Player) {
	s.playersMu.Lock()
	defer s.playersMu.Unlock()

	s.players[p.player] = p
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
=======
// Whitelist returns the server whitelist
func (server *Server) Whitelist() *list.List { return server.whitelist }
>>>>>>> Stashed changes
