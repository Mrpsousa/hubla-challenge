package database

import "github.com/mrpsousa/api/internal/entity"

type TransactionInterface interface {
	Create(transaction *entity.Transaction) error
	GetProductorBalance() ([]entity.DtoSellers, error)
	GetAssociateBalance() ([]entity.DtoSellers, error)
}
