package rest

import (
	data "lec-06/onClass/studentApi/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	k, _ := strconv.Atoi(key)
	db.Delete(&data.Student{}, k)
}
