package controller

import (
	"encoding/json"
	"fmt"
	"lec-06/exercise/managerStore/model"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	product = model.Product{}
)

func ReturnProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := model.ConnectDB()
	if err := db.Find(&model.Products).Error; err != nil {
		fmt.Fprintln(w, "failed to return all product items ")
	}
	for _, s := range model.Products {
		json.NewEncoder(w).Encode(&s)
	}
}

// ioutill.ReadAll to readall body of http request
// after that convert to struct type
// select  two fields productname and productprice which ensure user can not effect on id field
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&product)
	if err := db.Debug().Select("ProductName", "ProductPrice").Create(&product).Error; err != nil {
		fmt.Fprintf(w, "failed to update new item  with error : %v", err)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if err := db.First(&product, params["id"]); err != nil {
		fmt.Fprintf(w, "failed to update new item  with id := %v, with error : %v", params, err)
		return
	}
	json.NewDecoder(r.Body).Decode(&product)
	db.Save(&product)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := model.ConnectDB()
	vars := mux.Vars(r)
	params := vars["id"]
	fmt.Printf("Deleting at id = %v ", params)
	if err := db.Where("id = ?", params).Delete(&model.Product{}).Error; err != nil {
		fmt.Fprintf(w, "Get error while deleting at %v, with error:  %v", params, err)
		return
	}
}
