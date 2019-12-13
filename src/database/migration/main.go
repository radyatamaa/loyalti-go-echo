package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"

)


func main() {


	//guids := guid.New()
	//fmt.Println(guids)
	db := database.ConnectionDB()

}


