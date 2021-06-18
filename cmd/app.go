package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	Info   AppInfo
}

type AppInfo struct {
	Name        string
	Description string
	URL         string
}

func NewServer() *Server {
	s := Server{
		router: mux.NewRouter(),
		Info:   info,
	}

	s.setupRoutes()
	return &s
}

func (s *Server) Run() {
	serve := fmt.Sprintf("%s:%s", "0.0.0.0", cfg.Port)
	logger.Infow("Starting web service",
		"host", "0.0.0.0",
		"port", cfg.Port,
	)
	logger.Fatal(http.ListenAndServe(serve, s.router))
}
