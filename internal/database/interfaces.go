package database

import (
	"github.com/guiyones/brimos/internal/entity"
)

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll() ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
