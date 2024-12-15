package app

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/yikuanzz/social/internal/base/log"
	"github.com/yikuanzz/social/internal/base/server"
)

type Application struct {
	id      string
	name    string
	version string
	servers []server.Server
	signals []os.Signal
}

type Option func(application *Application)

// NewApp creates a new application instance
func NewApp(ops ...Option) *Application {
	app := &Application{}
	for _, op := range ops {
		op(app)
	}

	// default random id
	if len(app.id) == 0 {
		bytes := make([]byte, 24)
		_, _ = rand.Read(bytes)
		app.id = hex.EncodeToString(bytes)
	}
	// default accept signals
	if len(app.signals) == 0 {
		app.signals = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	}
	return app

}

// WithID application add id
func WithID(id string) func(application *Application) {
	return func(application *Application) {
		application.id = id
	}
}

// WithName application add name
func WithName(name string) func(application *Application) {
	return func(application *Application) {
		application.name = name
	}
}

// WithVersion application add version
func WithVersion(version string) func(application *Application) {
	return func(application *Application) {
		application.version = version
	}
}

// WithServer application add server
func WithServer(servers ...server.Server) func(application *Application) {
	return func(application *Application) {
		application.servers = servers
	}
}

// WithSignals application add listen signals
func WithSignals(signals []os.Signal) func(application *Application) {
	return func(application *Application) {
		application.signals = signals
	}
}

// Run application start
func (app *Application) Run(ctx context.Context) error {
	if len(app.servers) == 0 {
		return nil
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, app.signals...)
	errCh := make(chan error, 1)

	for _, s := range app.servers {
		go func(server server.Server) {
			if err := server.Start(); err != nil {
				log.Errorf("failed to start server, err: %s", err)
				errCh <- err
			}
		}(s)
	}

	select {
	case err := <-errCh:
		_ = app.Stop()
		return err
	case <-ctx.Done():
		return app.Stop()
	case <-quit:
		return app.Stop()
	}
}

// Stop application stop
func (app *Application) Stop() error {
	wg := sync.WaitGroup{}
	for _, s := range app.servers {
		wg.Add(1)
		go func(srv server.Server) {
			defer wg.Done()
			if err := srv.Shutdown(); err != nil {
				log.Errorf("failed to shutdown server, err: %s", err)
			}
		}(s)
	}
	// Wait all server graceful shutdown
	wg.Wait()
	return nil
}
