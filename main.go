package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Weight float64

type Product struct {
	ID     string
	Name   string
	Price  float64
	Weight Weight
}

type Client struct {
	ID   string
	Name string
}

type Sale struct {
	ID                string
	Client            Client
	SaleDate          time.Time
	Products          []Product
	ProductQuantity   []float64
	ProductTotalPrice []float64
	Discount          float64
	Freight           float64
	SaleTotalPrice    float64
	Payment           Payment
}

type BankAccount struct {
	ID             string
	Name           string
	AccountBalance float64
}

type Payment struct {
	ID          string
	BankAccount BankAccount
	PaymentType string
	PayDate     time.Time
	DueDate     time.Time
	Status      bool
}

// type SaleProduct struct {
// 	Sale       Sale
// 	Product    Product
// 	Quantity   float64
// 	TotalPrice float64
// }

// type Sale struct {
// 	ID        string
// 	Products  []SaleProduct2
// 	Client    Client
// 	SaleDate  time.Time
// 	SalePrice float64
// }

func NewProduct(name string, price float64, weight Weight) *Product {
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
	err = InsertProduct(db, *product)
	if err != nil {
		panic(err)
	}

	product.Price = 12.0
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectOneProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("O Produduto: %v pesa %vgr e tem o valor de R$%.2f\n", p.Name, p.Weight, p.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, ps := range products {
		fmt.Printf("Produduto: %v peso: %vgr valor:R$%.2f\n", ps.Name, ps.Weight, ps.Price)
	}

	err = deleteProduct(db, product.ID)
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

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ? , price = ? , weight = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Weight, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price, weight from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price, &p.Weight)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price, weight from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Weight)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil

}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
