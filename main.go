package main

import (
	"log"
	"net/http"
	han "tusharhow/ecom/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/addproduct", han.AddProduct).Methods("POST")
	r.HandleFunc("/getproduct", han.GetProduct).Methods("GET")
	r.HandleFunc("/getproduct/{id}", han.GetProductByID).Methods("GET")
	r.HandleFunc("/updateproduct/{id}", han.UpdateProduct).Methods("PUT")
	r.HandleFunc("/deleteproduct/{id}", han.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/totalprice", han.TotalPrice).Methods("GET")
	r.HandleFunc("/addtocart/{id}", han.AddToCart).Methods("POST")


	han.Products = append(han.Products, han.Product{ID: "1", Name: "Laptop", Price: "100"})
	han.Products = append(han.Products, han.Product{ID: "2", Name: "Mobile", Price: "200"})
	han.Products = append(han.Products, han.Product{ID: "3", Name: "TV", Price: "300"})

	log.Fatal(http.ListenAndServe(":8080", r))
}
