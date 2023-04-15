package service

import (
	"context"
	"onelab/internal/model"
	"onelab/internal/repository"
)

type ITransactionService interface {
	CreateTransaction(ctx context.Context, transactionsCreate model.TransactionsCreate) error
	GetAll(ctx context.Context) ([]model.Transactions, error)
}

type TransactionService struct {
	Repo *repository.Manager
}

func NewTransactionService(repo *repository.Manager) *TransactionService {
	return &TransactionService{
		Repo: repo,
	}
}

func (s TransactionService) CreateTransaction(ctx context.Context, transactions model.TransactionsCreate) error {
	//TODO validation
	err := s.Repo.Transactions.CreateTransaction(ctx, transactions)
	return err
}

func (s TransactionService) GetAll(ctx context.Context) ([]model.Transactions, error) {
	trn, err := s.Repo.Transactions.GetAll(ctx)
	if err != nil {
		return []model.Transactions{}, err
	}
	return trn, nil
}
