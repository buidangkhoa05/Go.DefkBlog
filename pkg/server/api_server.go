package server

import (
	"golang.org/x/net/context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":8080"
	_defaultShutdownTimeout = 3 * time.Second
)

type ApiServer struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewApiServer(handler http.Handler, options ...Option) *ApiServer {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	apiServer := &ApiServer{
		server:          httpServer,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// For more option configs
	for _, option := range options {
		option(apiServer)
	}

	apiServer.start()

	return apiServer
}

func (httpServer *ApiServer) start() {
	go func() {
		httpServer.notify <- httpServer.server.ListenAndServe()
		close(httpServer.notify)
	}()
}

// Notify when have error was thrown
func (httpServer *ApiServer) Notify() chan error {
	return httpServer.notify
}

func (httpServer *ApiServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), httpServer.shutdownTimeout)
	defer cancel()
	return httpServer.server.Shutdown(ctx)
}

func (httpServer *ApiServer) Server() *http.Server {
	return httpServer.server
}
