package model

import "time"

type Transactions struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	SenderID   uint      `json:"sender" gorm:"not null"`
	ReceiverID uint      `json:"receiver" gorm:"not null"`
	Amount     uint      `json:"amount" gorm:"not null"`
	Date       time.Time `json:"transfer_date" gorm:"not null"`
}

type TransactionsCreate struct {
	SenderID   uint `json:"sender"`
	ReceiverID uint `json:"receiver"`
	Amount     uint `json:"amount"`
}
