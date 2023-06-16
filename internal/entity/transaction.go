package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidID      = errors.New("id_is_required_or_invalid")
	ErrInvalidType    = errors.New("type_is_required_or_invalid")
	ErrInvalidDate    = errors.New("date_is_required_or_invalid")
	ErrInvalidProduct = errors.New("product_is_required_or_invalid")
	ErrInvalidValue   = errors.New("value_is_required_or_invalid")
	ErrInvalidSeller  = errors.New("seller_is_required_or_invalid")
)

type Transaction struct {
	ID        string
	Type      int8
	CreatedAt string `json:"created_at"` //TODO:  Data - ISO Date + GMT
	Product   string
	Value     float64
	Seller    string
}

func NewTransaction(tp int, dt, product, seller string, val float64) (*Transaction, error) {

	transaction := &Transaction{
		ID:        uuid.New().String(),
		Type:      int8(tp),
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
		return ErrInvalidID
	}
	if !(t.Type >= 1 && t.Type <= 4) {
		return ErrInvalidType
	}
	if t.CreatedAt == "" {
		return ErrInvalidDate
	}
	if t.Product == "" {
		return ErrInvalidProduct
	}
	if t.Value < 1.0 {
		return ErrInvalidValue
	}
	if t.Seller == "" {
		return ErrInvalidSeller
	}
	return nil
}
