package rest

import (
	"encoding/json"
	"io/ioutil"
	data "lec-06/onClass/studentApi/data""
	"net/http"
)

var (
	student  = data.Student{}
	students = data.Students
)

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &student)
	db.Debug().Model(&student).Where("id = ?", student.Id).Updates(&data.Student{Name: student.Name, Number: student.Number, Graduated: student.Graduated})
	json.NewEncoder(w).Encode(student)
}

func UpdateStudents(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &students)
	for _, s := range students {
		db.Debug().Model(&data.Student{}).Where("id = ?", s.Id).Updates(&data.Student{Name: s.Name, Number: s.Number, Graduated: s.Graduated})
	}
}
