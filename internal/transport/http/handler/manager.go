package handler

import (
	"onelab/internal/config"
	"onelab/internal/service"
)

type Manager struct {
	srv *service.Manager
}

func NewManager(cfg *config.Config, srv *service.Manager) *Manager {
	return &Manager{
		srv: srv,
	}
}
