package handler

import (
	"fmt"
	"github.com/lcok/vault-sync/internal/config"
	"github.com/lcok/vault-sync/internal/logger"
)

type Server struct {
	c   *config.EnvConfig
	log *logger.Logger
}

func NewServer(c *config.EnvConfig, log *logger.Logger) *Server {
	return &Server{
		c:   c,
		log: log,
	}
}

func (s *Server) Start() {
	// TODO impl this
	fmt.Println("Server is running...")
}
