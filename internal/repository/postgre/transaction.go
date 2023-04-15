package postgre

import (
	"context"
	"errors"
	"onelab/internal/model"

	"gorm.io/gorm"
)

type TransactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{
		DB: db,
	}
}

func (r *TransactionRepo) CreateTransaction(ctx context.Context, transactions model.TransactionsCreate) error {
	var senderBalance uint

	// Check if sender has enough balance
	if err := r.DB.Table("users").Where("id = ?", transactions.SenderID).Select("balance").Scan(&senderBalance).Error; err != nil {
		return err
	}
	if senderBalance < transactions.Amount {
		return errors.New("insufficient balance")
	}

	// Start transaction
	transactionsDB := r.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			transactionsDB.Rollback()
		}
	}()

	if err := transactionsDB.Table("transactions").Create(&transactions).Error; err != nil {
		transactionsDB.Rollback()
		return err
	}

	// Update sender balance
	if err := transactionsDB.Table("users").Where("id = ?", transactions.SenderID).Update("balance", gorm.Expr("balance - ?", transactions.Amount)).Error; err != nil {
		transactionsDB.Rollback()
		return err
	}

	// Update receiver balance
	if err := transactionsDB.Table("users").Where("id = ?", transactions.ReceiverID).Update("balance", gorm.Expr("balance + ?", transactions.Amount)).Error; err != nil {
		transactionsDB.Rollback()
		return err
	}

	// Commit transaction
	if err := transactionsDB.Commit().Error; err != nil {
		transactionsDB.Rollback()
		return err
	}

	return nil
}

func (r *TransactionRepo) GetAll(ctx context.Context) ([]model.Transactions, error) {
	var transactions []model.Transactions
	if err := r.DB.Table("transactions").Find(&transactions).Error; err != nil {
		return []model.Transactions{}, err
	}
	return transactions, nil
}
