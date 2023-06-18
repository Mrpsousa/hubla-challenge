package database

import (
	"fmt"

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

func (t *Transaction) ListProductorBalance() ([]entity.DtoQueryResult, error) {
	var associates []entity.DtoQueryResult
	var associate entity.DtoQueryResult
	rows, err := t.DB.Model(&entity.DtoQueryResult{}).Raw(`
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
		return nil, err // fazer tratamento errro <----------------
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&associate.Seller, &associate.TValue)
		if err != nil {
			return nil, err // fazer tratamento errro <----------------
		}
		associates = append(associates, associate)
		fmt.Println(associate.Seller, associate.TValue)
	}
	if err = rows.Err(); err != nil {
		return nil, err // fazer tratamento errro <----------------
	}

	return t.ConsolidateSellers(associates), nil
}

func (t *Transaction) ListAssociateBalance() ([]entity.DtoQueryResult, error) {
	var associates []entity.DtoQueryResult
	var associate entity.DtoQueryResult
	rows, err := t.DB.Model(&entity.DtoQueryResult{}).Raw(`
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
		return nil, err // fazer tratamento errro <----------------
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&associate.Seller, &associate.TValue)
		if err != nil {
			return nil, err // fazer tratamento errro <----------------
		}
		associates = append(associates, associate)
	}
	if err = rows.Err(); err != nil {
		return nil, err // fazer tratamento errro <----------------
	}

	return associates, nil
}

func (t *Transaction) ConsolidateSellers(dtoList []entity.DtoQueryResult) []entity.DtoQueryResult {
	consolidatedMap := make(map[string]float64)

	for _, dto := range dtoList {
		consolidatedMap[dto.Seller] += dto.TValue
	}

	consolidatedList := make([]entity.DtoQueryResult, 0)

	for seller, value := range consolidatedMap {
		consolidatedList = append(consolidatedList, entity.DtoQueryResult{
			Seller: seller,
			TValue: value,
		})
	}

	return consolidatedList
}
