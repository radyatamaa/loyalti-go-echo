package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
)
func GetCategory(page int, size int) []model.MerchantCategory {
	db:= database.ConnectionDB()
	var category []model.MerchantCategory
	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: page,
		Limit:	size,
		OrderBy:	[]string{"category_name desc"},
	}, &category)

	db.Close()
	return category
}
