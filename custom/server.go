package custom

import (
	"github.com/RestartFU/gophig"
	"github.com/RestartFU/list"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-plus/worldmanager"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
)

// Server embeds *server.Server.
type Server struct {
	*server.Server
	logger *logrus.Logger

	whitelist    *list.List
	worldManager *worldmanager.WorldManager

	conns   map[net.Addr]*minecraft.Conn
	connsMu sync.RWMutex

	allowedData map[string]*minecraft.Conn
}

// NewServer returns a new *Server
func NewServer(config *server.Config, logger *logrus.Logger) *Server {
	s := server.New(config, logger)
	wl, err := list.New(&list.Settings{CacheOnly: false, Gophig: gophig.NewGophig("./whitelist", "toml", 0777)})
	if err != nil {
		panic(err)
	}
	return &Server{
		Server:       s,
		logger:       logger,
		whitelist:    wl,
		worldManager: worldmanager.New(s, logger),
		
		conns:       make(map[net.Addr]*minecraft.Conn),
		allowedData: make(map[string]*minecraft.Conn),
	}
}

// Accept accepts the incoming connections and returns a new *custom.Player.
// It only returns an error when the server is closed.
func (server *Server) Accept() (*Player, error) {
	p, err := server.Server.Accept()
	if err != nil {
		return nil, err
	}
	conn, ok := server.Conn(p.Addr())
	if !ok {
		server.logger.Errorf("conn for player %s not found in map", p.Name())
	}
	return NewPlayer(p, conn), nil
}

// Whitelist returns the server whitelist
func (server *Server) Whitelist() *list.List { return server.whitelist }

// WorldManager will return the *worldmanager.WorldManager of the server
func (server *Server) WorldManager() *worldmanager.WorldManager { return server.worldManager }

// Conn returns the *minecraft.Conn with the passed net.Addr
func (server *Server) Conn(addr net.Addr) (*minecraft.Conn, bool) {
	server.connsMu.Lock()
	conn, ok := server.conns[addr]
	server.connsMu.Unlock()
	if ok {
		server.connsMu.RLock()
		delete(server.conns, addr)
		server.connsMu.RLock()
	}
	return conn, ok
}

// SetConn sets a *minecraft.Conn to the *server.conns map
func (server *Server) SetConn(conn *minecraft.Conn, addr net.Addr) {
	server.connsMu.Lock()
	defer server.connsMu.Unlock()
	server.conns[addr] = conn
}

// AllowedData ...
func (server *Server) AllowedData() map[string]*minecraft.Conn { return server.allowedData }

// SetAllowedData ...
func (server *Server) SetAllowedData(model string, conn *minecraft.Conn) {
	server.allowedData[model] = conn
}
