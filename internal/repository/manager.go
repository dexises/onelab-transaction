package repository

import (
	"context"

	"onelab/internal/config"
	"onelab/internal/model"
	"onelab/internal/repository/postgre"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.UserCreate) error
	Get(ctx context.Context, id int) (model.User, error)
}

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, transactionsCreate model.TransactionsCreate) error
	GetAll(ctx context.Context) ([]model.Transactions, error)
}

type Manager struct {
	pg           *gorm.DB
	User         IUserRepository
	Transactions ITransactionRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Manager, error) {
	db, err := postgre.NewConnection(cfg.DB)
	if err != nil {
		return nil, err
	}

	return &Manager{
		pg:           db,
		User:         postgre.NewUserRepo(db),
		Transactions: postgre.NewTransactionRepo(db),
	}, nil
}
