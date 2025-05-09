package entity

import (
	"errors"

	"github.com/guiyones/brimos/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID    entity.ID `json:"id"`
	Name  string    `json:"name"` //description
	Price float64   `json:"price"`

	// category
	// subCategory
	// product code
	// unit
	// NCM
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:    entity.NewID(),
		Name:  name,
		Price: price,
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price == 0.0 {
		return ErrInvalidPrice
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
