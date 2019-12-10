package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
)

func main() {

	db := database.ConnectionDB()

	//db.AutoMigrate(model.Card{})
	fmt.Println("migrasi berhasil")
}
