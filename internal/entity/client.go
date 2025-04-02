package entity

import "github.com/guiyones/brimos/pkg/entity"

type Client struct {
	ID              entity.ID `json:"id"`
	Name            string    `json:"name"`
	CorporateReason string    `json:"corporate-reason"`
	Cnpj            string    `json:"cnpj"`
	IE              string    `json:"ie"`

	PriceList PriceList `json:"price-list"`
	Address   Address   `json:"address"`
}

type Address struct {
	CEP      string `json:"cep"`
	City     string `json:"city"`
	State    string `json:"state"`
	District string `json:"district"`
	Street   string `json:"street"`
	Number   string `json:"number"`
}

type PriceList struct {
	ID       entity.ID
	Name     string
	Discount float64
}
