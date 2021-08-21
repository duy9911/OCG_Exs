package rest

import (
	"encoding/json"
	data "lec-05/Exercise/data"
	"net/http"
	"strconv"

	"lec-05/Exercise/handler"

	"github.com/gorilla/mux"
)

var (
	db, _ = data.ConnectDB()
)

func ReturnSingleStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	k, _ := strconv.Atoi(key)

	rows := db.First(&data.Students, k)
	handler.HandleError(rows.Error)

	rows.Scan(&data.Students)
	json.NewEncoder(w).Encode(&data.Students)
}

func ReturnStudents(w http.ResponseWriter, r *http.Request) {
	rows := db.Find(&data.Students)
	handler.HandleError(rows.Error)
	rows.Scan(&data.Students)

	for _, s := range data.Students {
		json.NewEncoder(w).Encode(&s)
	}

}
