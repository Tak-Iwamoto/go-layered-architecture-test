package infrastructure

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
