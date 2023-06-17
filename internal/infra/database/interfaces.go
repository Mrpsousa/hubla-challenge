package database

import "github.com/mrpsousa/api/internal/entity"

type TransactionInterface interface {
	Create(transaction *entity.Transaction) error
}
