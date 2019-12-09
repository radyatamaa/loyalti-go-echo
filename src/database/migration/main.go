package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func main() {

	db := database.ConnectionDB()

	//db.Model(&model.Program{}).AddIndex("category", "program_gallery")
	db.AutoMigrate(&model.Program{})

}
