package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/openware/rails5session-go"
	"github.com/openware/rangda/pkg/log"
)

const (
	CookieSession = "_barong_session"
)

type Server struct {
	chi.Router
	encryption *rails5session.Encryption
}

func NewServer(config *Config) *Server {
	encryption := rails5session.NewEncryption(
		[]byte(config.Session.SecretKeyBase),
		rails5session.DefaultEncryptedCookieSalt,
		rails5session.DefaultEncryptedSignedCookieSalt,
	)

	return &Server{
		// Router is an interface, need to create a *chi.Mux
		Router: chi.NewRouter(),

		encryption: encryption,
	}
}

func (server *Server) SetupRoutes() {
	server.Use(middleware.RequestID)
	server.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger: log.Logger,
	}))

	server.Get("/api/v1/auth", server.HandleAuth)
}
