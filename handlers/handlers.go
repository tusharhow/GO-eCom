package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

var Products []Product

// add product
func AddProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&Products)

	w.WriteHeader(http.StatusOK)

	d := json.NewEncoder(w).Encode(append(Products))
	fmt.Println(d)
}

// get product
func GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(append(Products))
}

// get product by id
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

// update product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Products {
		if item.ID == params["id"] {
			Products = append(Products[:index], Products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			Products = append(Products, product)
			json.NewEncoder(w).Encode(Products)
			return
		}
	}
	json.NewEncoder(w).Encode(Products)
}

// delete product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Products {
		if item.ID == params["id"] {
			Products = append(Products[:index], Products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Products)
}

// total price
func TotalPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var totalPrice int
	for _, item := range Products {
		price, _ := strconv.Atoi(item.Price)
		totalPrice += price
		if totalPrice != 0 {
			w.Write([]byte("Total price is: " + strconv.Itoa(totalPrice)))

		}
		w.Write([]byte("Cart is empty\n"))

	}

}

// add to cart
func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range Products {
		if item.ID == params["id"] {

			w.Write([]byte("Added to cart"))
		}
		w.Write([]byte("Product not found"))
	}

}
