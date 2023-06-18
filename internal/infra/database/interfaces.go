package database

import "github.com/mrpsousa/api/internal/entity"

type TransactionInterface interface {
	Create(transaction *entity.Transaction) error
	ListProductorBalance() ([]entity.DtoQueryResult, error)
	ListAssociateBalance() ([]entity.DtoQueryResult, error)
}
