package database

import (
	"errors"

	"github.com/mrpsousa/api/internal/entity"
	"gorm.io/gorm"
)

var (
	ErrorQueryListProducer  = errors.New("fail_to_query_producers")
	ErrorQueryListAssociate = errors.New("fail_to_query_associates")
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

func (t *Transaction) GetProductorBalance() ([]entity.DtoSellers, error) {
	var producers = make([]entity.DtoSellers, 0)
	var produtor entity.DtoSellers

	rows, err := t.DB.Model(&entity.DtoSellers{}).Raw(`
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

func (t *Transaction) GetAssociateBalance() ([]entity.DtoSellers, error) {
	var associates = make([]entity.DtoSellers, 0)
	var associate entity.DtoSellers

	rows, err := t.DB.Model(&entity.DtoSellers{}).Raw(`
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

func (t *Transaction) ConsolidateSellers(dtoList []entity.DtoSellers) []entity.DtoSellers {
	consolidatedMap := make(map[string]float64)

	for _, dto := range dtoList {
		consolidatedMap[dto.Seller] += dto.TValue
	}

	consolidatedList := make([]entity.DtoSellers, 0)

	for seller, value := range consolidatedMap {
		consolidatedList = append(consolidatedList, entity.DtoSellers{
			Seller: seller,
			TValue: value,
		})
	}

	return consolidatedList
}
