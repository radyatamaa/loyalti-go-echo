package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"

	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	//"net/http"
)

func CreateMerchant(merchant *model.Merchant) string{
	db := database.ConnectionDB()

	merchantObj := *merchant

	db.Create(&merchantObj)

	return merchantObj.MerchantEmail
}

func UpdateMerchant  (merchant *model.Merchant) string {
	db := database.ConnectionDB()
	db.Model(&merchant).Where("merchant_email = ?", merchant.MerchantEmail).Update(&merchant)
	return merchant.MerchantEmail
}

func DeleteMerchant(merchant *model.Merchant) string {
db:= database.ConnectionDB()
db.Model(&merchant).Where("merchant_email= ?",merchant.MerchantEmail).Delete(&merchant)
return "berhasil dihapus"
}


func GetMerchant(page *int, size *int, sort *int, email *string) []model.Merchant {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var merchant []model.Merchant
	db.Find(&merchant)

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil && email != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name desc"},
				}, &merchant)
				db.Order("merchant_name desc").Where("merchant_email = ?", email).Find(&merchant)
			}
		case 2:
			if size != nil && page != nil && email != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name asc"},
				}, &merchant)
				db.Order("merchant_name asc").Where("merchant_email = ?", email).Find(&merchant)
			}
		}
	}
	db.Close()
	return merchant
}
