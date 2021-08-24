package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lec-06/exercise/managerStore/model"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	customer = model.Customer{}
)

func ReturnCustomer(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	customers := model.Customers
	if err := db.Find(&customers).Error; err != nil {
		fmt.Fprintln(w, "failed to return all  customer items ")
	}
	for _, s := range customers {
		json.NewEncoder(w).Encode(&s)
	}
}

// ioutill.ReadAll to readall body of http request
// after that convert to struct type
// select  two fields only which ensure user can not effect on id field
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &customer)

	if err := db.Debug().Select("CustomerName", "CustomerInfo").Create(&customer).Error; err != nil {
		fmt.Fprintf(w, "failed to update new item  with error : %v", err)
		return
	}
	json.NewEncoder(w).Encode(customer)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	params := vars["id"]
	if err := db.First(&customer, params).Error; err != nil {
		fmt.Fprintf(w, "failed to update new item  with id := %v, with error : %v", params, err)
		return
	}
	json.NewDecoder(r.Body).Decode(&customer)
	db.Debug().
		Save(&customer)
	json.NewEncoder(w).Encode(customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	vars := mux.Vars(r)
	params := vars["id"]
	fmt.Printf("Deleting at id = %v ", params)
	if err := db.Debug().Where("id = ?", params).Delete(&customer).Error; err != nil {
		fmt.Fprintf(w, "Get error while deleting at %v, with error:  %v", params, err)
		return
	}
}
