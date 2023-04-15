package model

type User struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"not null;unique"`
	Balance uint   `json:"balance"`
}

type UserCreate struct {
	Name    string `json:"name"`
	Balance uint   `json:"balance"`
}
