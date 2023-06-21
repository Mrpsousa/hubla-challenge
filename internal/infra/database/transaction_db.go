package database

import (
	"errors"

	"github.com/mrpsousa/api/internal/dto"
	"github.com/mrpsousa/api/internal/entity"
	"gorm.io/gorm"
)

var (
	ErrorQueryListProducer       = errors.New("fail_to_query_producers")
	ErrorQueryListAssociate      = errors.New("fail_to_query_associates")
	ErrorQueryListForeignCourses = errors.New("fail_to_query_foreign_courses")
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

func (t *Transaction) GetProductorBalance() ([]dto.DtoSellers, error) {
	var producers = make([]dto.DtoSellers, 0)
	var produtor dto.DtoSellers

	rows, err := t.DB.Model(&dto.DtoSellers{}).Raw(`
	SELECT DISTINCT
    	t.seller,
    	s.tvalue
	FROM
    	transactions t
	JOIN (
    	SELECT
        	product,
       		SUM(
            	CASE
					WHEN type = 1 THEN value
					WHEN type = 2 THEN value
					WHEN type = 4 THEN -value
                	ELSE 0
            	END
        	) AS tvalue
    	FROM
        	transactions
    	GROUP BY
        	product
	) s ON t.product = s.product
	WHERE
    	t.type = 1;`).Rows()

	if err != nil {
		return nil, ErrorQueryListProducer

	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&produtor.Seller, &produtor.TValue)
		if err != nil {
			return nil, ErrorQueryListProducer
		}
		producers = append(producers, produtor)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrorQueryListProducer
	}

	return t.ConsolidateSellers(producers), nil
}

func (t *Transaction) GetAssociateBalance() ([]dto.DtoSellers, error) {
	var associates = make([]dto.DtoSellers, 0)
	var associate dto.DtoSellers

	rows, err := t.DB.Model(&dto.DtoSellers{}).Raw(`
	SELECT
    	seller,
    	SUM(value) AS tvalue
	FROM
    	transactions
	WHERE
    	type = 4
	GROUP BY
    	seller;`).Rows()

	if err != nil {
		return nil, ErrorQueryListAssociate
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&associate.Seller, &associate.TValue)
		if err != nil {
			return nil, ErrorQueryListAssociate
		}
		associates = append(associates, associate)
	}
	if err = rows.Err(); err != nil {
		return nil, ErrorQueryListAssociate
	}

	return associates, nil
}

func (t *Transaction) ConsolidateSellers(dtoList []dto.DtoSellers) []dto.DtoSellers {
	consolidatedMap := make(map[string]float64)

	for _, dto := range dtoList {
		consolidatedMap[dto.Seller] += dto.TValue
	}

	consolidatedList := make([]dto.DtoSellers, 0)

	for seller, value := range consolidatedMap {
		consolidatedList = append(consolidatedList, dto.DtoSellers{
			Seller: seller,
			TValue: value,
		})
	}

	return consolidatedList
}

func (t *Transaction) GetForeignCourses() ([]dto.DtoCourses, error) {
	var transactions []entity.Transaction
	var courses = make([]dto.DtoCourses, 0)

	err := t.DB.Where("foreign_product = ?", true).Group("seller").Find(&transactions).Error
	if err != nil {
		return nil, ErrorQueryListForeignCourses
	}
	for _, transaction := range transactions {
		dtoCourse := dto.DtoCourses{
			Type:      transaction.Type,
			CreatedAt: transaction.CreatedAt,
			Product:   transaction.Product,
			Value:     transaction.Value,
			Seller:    transaction.Seller,
		}
		courses = append(courses, dtoCourse)
	}

	return courses, nil
}
