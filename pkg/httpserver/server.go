package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	server *http.Server
	notify chan error
	Router *gin.Engine

	shutdownTimeout time.Duration
}

func New(opts ...Option) *Server {
	router := gin.Default()

	srv := &http.Server{
		Handler: router.Handler(),
	}

	for _, opt := range opts {
		opt(srv)
	}

	return &Server{
		server:          srv,
		notify:          make(chan error, 1),
		Router:          router,
		shutdownTimeout: _defaultShutdownTimeout,
	}
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
