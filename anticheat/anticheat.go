package anticheat

import (
	"errors"
	"github.com/RestartFU/practice/custom"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/justtaldevelops/oomph/player"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	LocalAddress  string
	RemoteAddress string
}

type AntiCheat struct {
	server *custom.Server

	config *Config

	logger *logrus.Logger
}

func New(config *Config, server *custom.Server, logger *logrus.Logger) *AntiCheat {
	return &AntiCheat{
		server: server,
		config: config,
		logger: logger,
	}
}

func (ac *AntiCheat) Start() {
	config := ac.config
	p, err := minecraft.NewForeignStatusProvider(config.RemoteAddress)
	if err != nil {
		panic(err)
	}
	listener, err := minecraft.ListenConfig{
		StatusProvider: p,
	}.Listen("raknet", config.LocalAddress)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	for {
		c, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go ac.handleConn(c.(*minecraft.Conn), listener, config)
	}
}

// handleConn handles a new incoming minecraft.Conn from the minecraft.Listener passed.
func (ac *AntiCheat) handleConn(conn *minecraft.Conn, listener *minecraft.Listener, config *Config) {
	src := rand.NewSource(time.Now().Unix())

	clientData := conn.ClientData()
	clientData.DeviceModel = strconv.Itoa(int(src.Int63()))

	ac.server.SetAllowedData(clientData.DeviceModel, conn)
	
	serverConn, err := minecraft.Dialer{
		EnableClientCache: true,
		IdentityData:      conn.IdentityData(),
		ClientData:        clientData,
	}.Dial("raknet", config.RemoteAddress)
	if err != nil {
		listener.Disconnect(conn, err.Error())
		return
	}

	var g sync.WaitGroup
	g.Add(2)
	go func() {
		if err := conn.StartGame(serverConn.GameData()); err != nil {
			return
		}
		g.Done()
	}()
	go func() {
		if err := serverConn.DoSpawn(); err != nil {
			return
		}
		g.Done()
	}()
	g.Wait()

	lg := ac.logger

	viewDistance := int32(8)
	p := player.NewPlayer(lg, world.Overworld, viewDistance, conn, serverConn)

	g.Add(2)
	go func() {
		defer listener.Disconnect(conn, "connection lost")
		defer serverConn.Close()
		for {
			pk, err := conn.ReadPacket()
			if err != nil {
				return
			}
			p.Process(pk, conn)
			if err := serverConn.WritePacket(pk); err != nil {
				if disconnect, ok := errors.Unwrap(err).(minecraft.DisconnectError); ok {
					_ = listener.Disconnect(conn, disconnect.Error())
				}
				return
			}
		}
		g.Done()
	}()
	go func() {
		defer serverConn.Close()
		defer listener.Disconnect(conn, "connection lost")
		for {
			pk, err := serverConn.ReadPacket()
			if err != nil {
				if disconnect, ok := errors.Unwrap(err).(minecraft.DisconnectError); ok {
					_ = listener.Disconnect(conn, disconnect.Error())
				}
				return
			}
			p.Process(pk, serverConn)
			if err := conn.WritePacket(pk); err != nil {
				return
			}
		}
		g.Done()
	}()
	g.Wait()
	p.Close()
}
