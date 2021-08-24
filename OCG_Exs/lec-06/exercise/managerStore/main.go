package main

import (
	"fmt"
	"lec-06/exercise/managerStore/controller"
	mid "lec-06/exercise/managerStore/middleware"
	"lec-06/exercise/managerStore/model"
	"log"
	"net/http"

	m "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	mRoute := m.NewRouter()
	mRoute.HandleFunc("/", homePage)
	mRoute.Use(mid.WithLogging)
	mRoute.HandleFunc("/api/customers", mid.BasicAuth(controller.CreateCustomer)).Methods("POST")
	mRoute.HandleFunc("/api/customers", mid.BasicAuth(controller.ReturnCustomer)).Methods("GET")
	mRoute.HandleFunc("/api/customers/{id}", mid.BasicAuth(controller.UpdateCustomer)).Methods("PUT")
	mRoute.HandleFunc("/api/customers/{id}", mid.BasicAuth(controller.DeleteCustomer)).Methods("DELETE")

	mRoute.HandleFunc("/api/products", mid.BasicAuth(controller.CreateProduct)).Methods("POST")
	mRoute.HandleFunc("/api/products", mid.BasicAuth(controller.ReturnProducts)).Methods("GET")
	mRoute.HandleFunc("/api/products/{id}", mid.BasicAuth(controller.UpdateProduct)).Methods("PUT")
	mRoute.HandleFunc("/api/products/{id}", mid.BasicAuth(controller.DeleteProduct)).Methods("DELETE")

	mRoute.HandleFunc("/api/orders", mid.BasicAuth(controller.CreateOrder)).Methods("POST")
	mRoute.HandleFunc("/api/orders", mid.BasicAuth(controller.ReturnOrders)).Methods("GET")
	mRoute.HandleFunc("/api/orders/{id}", mid.BasicAuth(controller.UpdateOrder)).Methods("PUT")
	mRoute.HandleFunc("/api/orders/{id}", mid.BasicAuth(controller.DeleteOrder)).Methods("DELETE")

	mRoute.HandleFunc("/api/orderdetails", mid.BasicAuth(controller.CreateOrderDetail)).Methods("POST")
	mRoute.HandleFunc("/api/orderdetails", mid.BasicAuth(controller.ReturnOrderDetails)).Methods("GET")
	mRoute.HandleFunc("/api/orderdetails/{id}", mid.BasicAuth(controller.UpdateOrderDetail)).Methods("PUT")
	mRoute.HandleFunc("/api/orderdetails/{id}", mid.BasicAuth(controller.DeleteOrderDetail)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", mRoute))

}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	model.CreateTables()
	handleRequests()
}
