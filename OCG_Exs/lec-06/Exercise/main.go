package main

import (
	"fmt"
	"lec-05/Exercise/rest"
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
	mRoute.HandleFunc("/articles", rest.ReturnStudents).Methods("GET")
	mRoute.HandleFunc("/article/{id}", rest.ReturnSingleStudent).Methods("GET")

	mRoute.HandleFunc("/article/{id}", rest.DeleteStudent).Methods("DELETE")

	mRoute.HandleFunc("/article", rest.CreateStudent).Methods("POST")
	mRoute.HandleFunc("/articles", rest.CreateStudents).Methods("POST")

	mRoute.HandleFunc("/article", rest.UpdateStudent).Methods("PUT")
	mRoute.HandleFunc("/articles", rest.UpdateStudents).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", mRoute))
}

func main() {
	//dummy data
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
