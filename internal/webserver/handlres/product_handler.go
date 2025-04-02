package handlres

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/guiyones/brimos/internal/database"
	"github.com/guiyones/brimos/internal/dto"
	"github.com/guiyones/brimos/internal/entity"

	entityPkg "github.com/guiyones/brimos/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct 	godoc
// @Summary 		Create product
// @Description 	Create products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			request 	body 	dto.CreateProductInput  true  "prduct request"
// @Success 		201
// @Failure 		500		{object} 	Error
// @Router 			/products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		msg := struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetOneProduct 	godoc
// @Summary 		Get a product
// @Description 	Get a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id 		path  		string  true  "product ID" Format(uuid)
// @Success 		200 	{object} 	entity.Product
// @Failure 		404
// @Failure 		500		{object} 	Error
// @Router 			/products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetOneProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("content-type", "apllication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.ProductDB.FindAll()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct 	godoc
// @Summary 		Update a product
// @Description 	Update a products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id 		path  		string  true  "product ID" Format(uuid)
// @Param 			request 	body 	dto.CreateProductInput  true  "product request"
// @Success 		200
// @Failure 		404
// @Failure 		500		{object} 	Error
// @Router 			/products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Para ver se o produto existe
	_, err = h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteProduct 	godoc
// @Summary 		Delete a product
// @Description 	Delete a products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id 		path  		string  true  "product ID" Format(uuid)
// @Success 		200
// @Failure 		404
// @Failure 		500		{object} 	Error
// @Router 			/products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
