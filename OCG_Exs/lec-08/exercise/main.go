package main

import (
	"exercise/controller"
	"exercise/model"
)

var db = model.ConnectDb()

var s []model.CsvLine

func main() {
	// controller.AddToSql()
	// controller.AddToEs()

	// db.Debug().Find(&s)

	// db.Debug().Where("Title LIKE ?", "%computer%").Find(&s)

	controller.QueryEs("Title", "gun")

}
