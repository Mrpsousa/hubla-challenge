package database

import (
	"github.com/mrpsousa/api/internal/entity"
	"gorm.io/gorm"
)

type Transaction struct {
	DB *gorm.DB
}

func NewTransaction(db *gorm.DB) *Transaction {
	return &Transaction{DB: db}
}

func (p *Transaction) Create(user *entity.Transaction) error {
	return p.DB.Create(user).Error
}
