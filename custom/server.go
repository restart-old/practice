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

// Whitelist returns the server whitelist
func (server *Server) Whitelist() *list.List { return server.whitelist }
