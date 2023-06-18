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

func (t *Transaction) ListProductorBalance() ([]entity.Producer, error) {
	var producers []entity.Producer
	var producer entity.Producer
	rows, err := t.DB.Model(&entity.Producer{}).Raw(`SELECT
	seller,
	SUM(
		CASE
			WHEN type = 1 THEN value
			WHEN type = 3 THEN -value
			ELSE 0
		END
	) AS tvalue
	FROM
		transactions
	GROUP BY
	seller`).Rows()

	if err != nil {
		return nil, err // fazer tratamento errro <----------------
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&producer.Seller, &producer.TValue)
		if err != nil {
			return nil, err // fazer tratamento errro <----------------
		}
		producers = append(producers, producer)
		fmt.Println(producer.Seller, producer.TValue)
	}
	if err = rows.Err(); err != nil {
		return nil, err // fazer tratamento errro <----------------
	}

	return producers, nil
}
