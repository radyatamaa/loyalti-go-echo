package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetMerchant(page *int, size *int, sort *int, email *string) []model.Merchant {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var merchant []model.Merchant
	db.Find(&merchant)

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name desc"},
				}, &merchant)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name asc"},
				}, &merchant)
			}
		}
	}

	if email != nil {
		if page != nil && size != nil {
			pagination.Paging(&pagination.Param{
				DB: 	db,
				Page: *page,
				Limit: *size,
				OrderBy: []string{"merchant_name asc"},

			}, &merchant)

			db.Order("id").Where("merchant_email = ?", email).Find(&merchant)
			//fmt.Println("2")
			//rows, err = db.Where("merchant_id = ?", id).Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			//if err != nil {
			//	panic(err)
			//}
		} else{
			db.Order("id").Where("merchant_email = ?", email).Find(&merchant)
		}
	}



	db.Close()
	return merchant
}
