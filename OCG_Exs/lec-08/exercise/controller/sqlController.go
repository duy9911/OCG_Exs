package controller

import (
	"fmt"

	"exercise/model"

	"gorm.io/gorm"
)

func AddToSql() {
	s := []model.CsvLine{}

	db := model.ConnectDb()

	lines, err := model.ReadCsv("./test.csv")
	if err != nil {
		panic(err)
	}

	db.Debug().Migrator().DropTable(&model.CsvLine{})
	db.AutoMigrate(&model.CsvLine{})

	// Loop through lines & turn into object
	for i, line := range lines {
		data := model.CsvLine{
			Type:  line[0],
			Title: line[1],
			Body:  line[2],
		}
		if i == 10000 {
			break
		}
		s = append(s, data)
		fmt.Println("Added row", i)
	}
	db = db.Session(&gorm.Session{CreateBatchSize: 1000})
	db.Create(s)

	fmt.Println("Done...!")
}
