package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetMerchant(page *int, size *int, email *string) []model.Merchant{
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var merchant []model.Merchant
	db.Find(&merchant)
	if size != nil && page != nil && email !=nil {
		pagination.Paging(&pagination.Param{
			DB:	db,
			Page: *page,
			Limit:	*size,
			OrderBy:	[]string{"merchant_name desc"},
		}, &merchant)
		db.Where("merchant_email = ?", email).Find(&merchant)
	}

	db.Close()
	return merchant
}
