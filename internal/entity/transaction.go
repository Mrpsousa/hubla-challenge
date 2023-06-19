package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/mrpsousa/api/pkg"
)

var base = "2006-01-02T15:04:05Z07:00"

type Transaction struct {
	ID             string
	Type           int8
	CreatedAt      string `json:"created_at"` //TODO:  Data - ISO Date + GMT
	Product        string
	Value          float64
	Seller         string
	ForeignProduct bool
}

func NewTransaction(tp int8, dt, product, seller string, val float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:             uuid.New().String(),
		Type:           tp,
		CreatedAt:      dt,
		Product:        product,
		Value:          (val / 100),
		Seller:         seller,
		ForeignProduct: pkg.ForeignProductValidate(product),
	}
	transaction.CreatedAt = transaction.DateConvert()

	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.ID == "" {
		return pkg.ErrInvalidID
	}
	if !(t.Type >= 1 && t.Type <= 4) {
		return pkg.ErrInvalidType
	}
	if t.CreatedAt == "" {
		return pkg.ErrInvalidDate
	}
	if t.Product == "" {
		return pkg.ErrInvalidProduct
	}
	if t.Value < 1.0 {
		return pkg.ErrInvalidValue
	}
	if t.Seller == "" {
		return pkg.ErrInvalidSeller
	}
	_, err := time.Parse(base, t.CreatedAt)
	if err != nil {
		return pkg.ErrInvalidDate
	}

	return nil
}

func (t *Transaction) DateConvert() string {
	if t.CreatedAt == "" {
		return ""
	}
	date, err := time.Parse(time.RFC3339, t.CreatedAt)
	if err != nil {
		return ""
	}
	return date.UTC().Format(base)
}
