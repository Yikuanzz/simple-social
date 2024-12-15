package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const DefaultShutdownTimeout = time.Minute

type GinServer struct {
	ShutdownTimeout time.Duration
	srv             *http.Server
}

type Options func(*GinServer)

// NewServer returns a new gin server with default settings
func NewServer(e *gin.Engine, addr string, options ...Options) *GinServer {
	server := GinServer{
		ShutdownTimeout: DefaultShutdownTimeout,
		srv: &http.Server{
			Addr:    addr,
			Handler: e,
		},
	}

	for _, option := range options {
		option(&server)
	}

	return &server
}

// WithShutdownTimeout duration of graceful shutdown
func WithShutdownTimeout(duration time.Duration) Options {
	return func(server *GinServer) {
		server.ShutdownTimeout = duration
	}
}

// Start to start the server and wait for it to listen on the given address
func (s *GinServer) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown shuts down the server and close with graceful shutdown duration
func (s *GinServer) Shutdown() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
