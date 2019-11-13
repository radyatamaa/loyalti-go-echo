package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
)

func GetAll(page int, size int) []model.Merchant{
	db := database.ConnectionDB()
	var merchant []model.Merchant
	db.Find(&merchant)
	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: page,
		Limit:	size,
		OrderBy:	[]string{"merchant_name desc"},
	}, &merchant)

	db.Close()
	return merchant
}
