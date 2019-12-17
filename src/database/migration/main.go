package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)


func main() {


	//guids := guid.New()
	//fmt.Println(guids)
	db := database.ConnectionDB()

	db.Model(&model.Card{}).DropColumn("tier_id")
}


