package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func NewPostgres() (*gorm.DB, error) {
	connectionString := "postgresql://postgres:12345678a@localhost:5432/postgres?sslmode=disable"
	DB, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	DB.DB().SetMaxIdleConns(2)
	DB.DB().SetMaxOpenConns(10)
	DB.DB().SetConnMaxLifetime(time.Second * 60)
	DB.LogMode(true)

	return DB, nil
}
