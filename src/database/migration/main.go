package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/database"

	"time"
)

func main() {
	var id = 16
	var pay = 7000
	repository.TotalPoint(&id, &pay)

	//guids := guid.New()
	//fmt.Println(guids)

)

func main() {

	db := database.ConnectionDB()
	//db.AutoMigrate(model.Card{})
	fmt.Println("migrasi berhasil")

}

