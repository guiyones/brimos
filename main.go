package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Weight int
}

func NewProduct(name string, price float64, weight int) *Product {
	return &Product{
		ID:     uuid.New().String(),
		Name:   name,
		Price:  price,
		Weight: weight,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/brimos")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Coalhada Sal", 6.20, 120)
	fmt.Println(product)
	err = InsertProduct(db, *product)

	if err != nil {
		panic(err)
	}
}

func InsertProduct(db *sql.DB, product Product) error {

	stmt, err := db.Prepare("insert into products(id, name, price, weight) values(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price, product.Weight)
	if err != nil {
		return err
	}

	return nil
}
