package database

import (
	"github.com/mrpsousa/api/internal/dto"
	"github.com/mrpsousa/api/internal/entity"
)

type TransactionInterface interface {
	Create(transaction *entity.Transaction) error
	GetProductorBalance() ([]dto.DtoSellers, error)
	GetAssociateBalance() ([]dto.DtoSellers, error)
	GetForeignCourses() ([]dto.DtoCourses, error)
}
