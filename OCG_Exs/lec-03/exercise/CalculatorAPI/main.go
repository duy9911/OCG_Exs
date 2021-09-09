package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	A      int64 `json:"a"`
	B      int64 `json:"b"`
	Result int64 `json:"result"`
}

func main() {
	r := mux.NewRouter()
	fmt.Print("Start")
	r.HandleFunc("/{math}/{a:[-1-9]+}/{b:[-1-9]+}", Calculator) //eg:localhost:{port}/*/2/3
	log.Fatal(http.ListenAndServe(":8000", r))                  //  endpoint
}
func Calculator(w http.ResponseWriter, r *http.Request) {
	math := mux.Vars(r)["math"]
	a := mux.Vars(r)["a"]
	b := mux.Vars(r)["b"]
	fmt.Print(math, a, b)
	result := Result{}
	result.A, _ = strconv.ParseInt(a, 10, 64)
	result.B, _ = strconv.ParseInt(b, 10, 64)

	switch math {
	case "+":
		result.Result = result.A + result.B
	case "-":
		result.Result = result.A - result.B
	case "*":
		result.Result = result.A * result.B
	case ":":
		result.Result = result.A / result.B
	default:
		fmt.Fprintf(w, "follow format :localhost:math/*/numberic_type/numberic_type")
	}

	bb, err := json.Marshal(result) // convert struct to json
	if err != nil {
		fmt.Println("error:", err)
	} // print on screen
	// w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin")) // set CORs to improve security
	// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(bb))
}
