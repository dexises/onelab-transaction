package postgre

import (
	"context"
	"database/sql"
	"errors"
	"onelab/internal/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: DB,
	}
}

func (r *UserRepo) Create(ctx context.Context, user model.UserCreate) error {
	return r.DB.WithContext(ctx).Table("users").Create(&user).Error
}

func (r *UserRepo) Get(ctx context.Context, id int) (model.User, error) {
	var user model.User
	if err := r.DB.Table("users").Unscoped().First(&user, id).Error; err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, sql.ErrNoRows
		}
		return model.User{}, err
	}
	return user, nil
}
