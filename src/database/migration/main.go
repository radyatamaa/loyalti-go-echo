package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func main() {

	db := database.ConnectionDB()

	db.AutoMigrate(model.Program{})
	fmt.Println("migrasi berhasil")

}
