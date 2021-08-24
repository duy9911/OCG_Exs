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
	order = model.Order{}
)

func ReturnOrders(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	orders := model.Orders
	if err := db.Find(&orders).Error; err != nil {
		fmt.Fprintln(w, "failed to return all  order items ")
	}
	for _, s := range orders {
		json.NewEncoder(w).Encode(&s)
	}
}

// ioutill.ReadAll to readall body of http request
// after that convert to struct type
// select  two fields only which ensure user can not effect on id field
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &order)

	if err := db.Debug().Select("CustomerID", "OrderDate", "PaymentMethod").Create(&order).Error; err != nil {
		fmt.Fprintf(w, "failed to create new item in order with id := %v, with error : %v", order.ID, err)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	params := vars["id"]
	if err := db.First(&order, params).Error; err != nil {
		fmt.Fprintf(w, "failed to update new item  with id := %v, with error : %v", params, err)
		return
	}
	json.NewDecoder(r.Body).Decode(&order)
	db.Debug().Save(&order)
	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	vars := mux.Vars(r)
	params := vars["id"]
	fmt.Printf("Deleting at id = %v ", params)
	if err := db.Debug().Where("id = ?", params).Delete(&model.Order{}).Error; err != nil {
		fmt.Fprintf(w, "Get error while deleting at %v, with error:  %v", params, err)
		return
	}
}
