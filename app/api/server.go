package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/tibeahx/growers-dairy/app/usecase"
)

type Server struct {
	httpServer *http.Server
	service    *usecase.ServiceProvider
}

func NewServer(service *usecase.ServiceProvider) *Server {
	return &Server{
		service: service,
	}
}

func (s *Server) Run(listenAddr string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         listenAddr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("server started on port %v", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

// todo: graceful shutdown
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
