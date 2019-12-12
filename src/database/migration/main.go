package main

import (
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"fmt"
)

func main() {
	//var id = 16
	//var pay = 7000
	//repository.TotalPoint(&id, &pay)

	//guids := guid.New()
	//fmt.Println(guids)

	db := database.ConnectionDB()
	db.AutoMigrate(&model.ProductCategory{})
	fmt.Println("migrasi berhasil")
}

