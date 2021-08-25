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
	orderdetail = model.OrderDetail{}
)

func ReturnOrderDetails(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	if err := db.Find(&model.OrderDetails).Error; err != nil {
		fmt.Fprintln(w, "failed to return all orderdetail items ", err)
	}
	for _, s := range model.OrderDetails {
		json.NewEncoder(w).Encode(&s)
	}
}

// ioutill.ReadAll to readall body of http request
// after that convert to struct type
// select  two fields productname and productprice which ensure user can not effect on id field
func CreateOrderDetail(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &orderdetail)

	if err := db.Debug().Select("ProductID", "OrderID", "Quantity").Create(&orderdetail).Error; err != nil {
		fmt.Fprintf(w, "failed to update new item  with error : %v", err)
		return
	}

	json.NewEncoder(w).Encode(orderdetail)
}

func UpdateOrderDetail(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	params := vars["id"]

	if err := db.First(&orderdetail, params).Error; err != nil {
		fmt.Fprintf(w, "failed to update new item  with id := %v, with error : %v", params, err)
		return
	}
	json.NewDecoder(r.Body).Decode(&orderdetail)
	db.Save(&orderdetail)
	json.NewEncoder(w).Encode(orderdetail)
}

func DeleteOrderDetail(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	vars := mux.Vars(r)
	params := vars["id"]
	fmt.Printf("Deleting at id = %v ", params)
	if err := db.Where("id = ?", params).Delete(&model.OrderDetail{}).Error; err != nil {
		fmt.Fprintf(w, "Get error while deleting at %v, with error:  %v", params, err)
		return
	}
}
