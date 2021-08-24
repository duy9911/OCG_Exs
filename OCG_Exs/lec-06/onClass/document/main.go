package main

import (
	"fmt"
	"lec-05/onClass/document/apis"
	"lec-05/onClass/document/middle"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/product", middle.BasicAuth(apis.FindAll)).Methods("GET")
	r.Handle("/product/search", middle.BasicAuth(apis.Search)).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Print(err)
	}
}
