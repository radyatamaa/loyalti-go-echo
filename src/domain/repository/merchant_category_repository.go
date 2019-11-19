package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
)

func GetCategory(page int, size int) []model.MerchantCategory {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var category []model.MerchantCategory
	db.Find(&category)

	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: page,
		Limit:	size,
		OrderBy:	[]string{"category_name desc"},
	}, &category)

	db.Close()
	return category
}