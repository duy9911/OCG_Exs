package rest

import (
	"encoding/json"
	"io/ioutil"
	"lec-06/onClass/studentApi/data"
	"lec-06/onClass/studentApi/handler"
	"net/http"
)

func CreateStudents(w http.ResponseWriter, r *http.Request) {
	db, err := data.ConnectDB()
	handler.HandleError(err)

	reqBody, _ := ioutil.ReadAll(r.Body)
	students := data.Students

	json.Unmarshal(reqBody, &students)
	db.Debug().Select("Name", "Number", "Graduated").Create(&students)

	for _, s := range students {
		json.NewEncoder(w).Encode(s)
	}
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	db, err := data.ConnectDB()
	handler.HandleError(err)

	reqBody, err2 := ioutil.ReadAll(r.Body)
	handler.HandleError(err2)

	student := data.Student{}
	json.Unmarshal(reqBody, &student)
	db.Debug().Select("Name", "Number", "Graduated").Create(&student)
}
