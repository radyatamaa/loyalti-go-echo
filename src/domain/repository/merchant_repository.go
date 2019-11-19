package repository

import (

)

func GetAll(page int, size int, email string) []model.Merchant{
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var merchant []model.Merchant
	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: page,
		Limit:	size,
		OrderBy:	[]string{"merchant_name desc"},
	}, &merchant)
	db.Where("merchant_email = ?", email).Find(&merchant)

	db.Close()
	return merchant
}
