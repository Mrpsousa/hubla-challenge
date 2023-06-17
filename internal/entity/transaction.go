package entity

import (
	"github.com/google/uuid"
	"github.com/mrpsousa/api/pkg"
)

type Transaction struct {
	ID        string
	Type      int8
	CreatedAt string `json:"created_at"` //TODO:  Data - ISO Date + GMT
	Product   string
	Value     float64
	Seller    string
}

func NewTransaction(tp int8, dt, product, seller string, val float64) (*Transaction, error) {

	transaction := &Transaction{
		ID:        uuid.New().String(),
		Type:      tp,
		CreatedAt: dt,
		Product:   product,
		Value:     val,
		Seller:    seller,
	}
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
	return nil
}
