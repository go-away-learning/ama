package server

import (
	"net/http"
	"time"

	"github.com/go-away-learning/ama/internal/database"
)

type Server struct {
	addr    string
	queries *database.Queries
}

func New(addr string, queries *database.Queries) *http.Server {
	NewServer := &Server{
		addr:    addr,
		queries: queries,
	}

	return &http.Server{
		Addr:         NewServer.addr,
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}
