package postgre

import (
	"fmt"

	"onelab/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg config.DBConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.Sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
