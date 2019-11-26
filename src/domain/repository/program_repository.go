package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetProgram(page *int, size *int,sort *int, category *int) []model.Program {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var program []model.Program
	db.Find(&program)
}