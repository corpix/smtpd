package smtpd

import (
	"fmt"
	"net"
	"os"
)

// Server describes a general data structures required to start SMTP server.
type Server struct {
	config  *ServerConfig
	pool    chan net.Conn
	handler Handler
}

// ServerConfig is a configuration for Server.
type ServerConfig struct {
	PoolSize       int
	ProcessThreads int
}

// Handler is a SMTP message(Envelope) handler.
type Handler interface {
	ServeSMTP(net.Conn, *Envelope)
}

// DefaultServerConfig is a default values for Server to start with.
var DefaultServerConfig = &ServerConfig{
	PoolSize:       16,
	ProcessThreads: 4,
}

// Serve serves SMTP protocol for Listener.
func (s *Server) Serve(l net.Listener) error {
	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		s.pool <- c
	}
}

// invokeProcessLoops starts as much ProcessThreads as ServerConfig require.
func (s *Server) invokeProcessLoops() {
	for i := 0; i < s.config.ProcessThreads; i++ {
		go s.processLoop()
	}
}

// processLoop reads a connection from pool and send it to processConn.
func (s *Server) processLoop() {
	for c := range s.pool {
		s.processConn(c)
	}
}

// processConn handles a connection and parsing the data according to the SMTP protocol.
func (s *Server) processConn(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error while handling SMTP connection: %s\n", err)
			return
		}
	}()
	defer conn.Close()

	smtpConn := NewConn(conn, Config{}, os.Stderr)
	var (
		event    Event
		envelope *Envelope
	)

	envelope = &Envelope{}

	for {
		event = s.processEvent(smtpConn.Next(), envelope)
		if event == DONE || event == ABORT {
			s.handler.ServeSMTP(conn, envelope)
			envelope = &Envelope{}
			break
		}
	}
}

// processEvent handles SMTP event filling Envelope with data.
func (s *Server) processEvent(event EventInfo, envelope *Envelope) Event {
	switch event.What {
	case COMMAND:
		switch event.Cmd {
		case MAILFROM:
			envelope.From = event.Arg
		case RCPTTO:
			envelope.To = append(envelope.To, event.Arg)
		}
	case GOTDATA:
		envelope.Data = []byte(event.Arg)
	}

	return event.What
}

// Serve serves SMTP protocol with default configuration.
func Serve(l net.Listener, h Handler) error {
	srv, err := New(nil, h)
	if err != nil {
		return err
	}

	return srv.Serve(l)
}

// New creates a new Server with specified ServerConfig(pass nil to use the default) and Handler.
func New(cfg *ServerConfig, handler Handler) (*Server, error) {
	if cfg == nil {
		cfg = DefaultServerConfig
	}

	pool := make(chan net.Conn, cfg.PoolSize)
	srv := &Server{
		config:  cfg,
		pool:    pool,
		handler: handler,
	}
	srv.invokeProcessLoops()

	return srv, nil
}
