package service

import (
	"onelab/internal/config"
	"onelab/internal/repository"
)

type Manager struct {
	User         IUserService
	Transactions ITransactionService
}

func NewService(cfg *config.Config, repo *repository.Manager) *Manager {
	return &Manager{
		User:         NewUserService(repo),
		Transactions: NewTransactionService(repo),
	}
}
