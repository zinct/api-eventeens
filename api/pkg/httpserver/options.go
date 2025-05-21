package httpserver

import (
	"net"
	"net/http"
	"time"
)

// Option -.
type Option func(*http.Server)

// Port -.
func Port(port string) Option {
	return func(s *http.Server) {
		s.Addr = net.JoinHostPort("", port)
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *http.Server) {
		s.ReadTimeout = timeout
		s.WriteTimeout = timeout
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *http.Server) {
		s.ReadTimeout = timeout
	}
}
