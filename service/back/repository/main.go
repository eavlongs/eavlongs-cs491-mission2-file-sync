package repository

import "gorm.io/gorm"

type MainRepository struct {
	db *gorm.DB
}

func NewMainRepository(db *gorm.DB) *MainRepository {
	return &MainRepository{db: db}
}
