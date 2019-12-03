package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
//	db := database.ConnectionDB()
//	//db := database.ConnectPostgre()
//	//connect postgre
//	db.AutoMigrate(&model.Merchant{})
//	merchant := &model.Merchant{}
//	repository.DeleteMerchant(merchant, "contact@starbucks.com")
//
//	merchant = &model.Merchant{
//		Created:               time.Now(),
//		CreatedBy:             "Admin",
//		Modified:              time.Now(),
//		ModifiedBy:            "Admin",
//		Active:                true,
//		IsDeleted:             false,
//		Deleted:               nil,
//		Deleted_by:            "",
//		MerchantName:          "Gojek",
//		MerchantEmail:         "contact@Gojek.com",
//		MerchantPhoneNumber:   "0819820201112",
//		MerchantProvince:      "DKI Jakarta",
//		MerchantCity:          "Jakarta",
//		MerchantAddress:       "Blok M",
//		MerchantPostalCode:    "19880",
//		MerchantCategoryId:    "1",
//		MerchantWebsite:       "www.gojekindonesia.com",
//		MerchantMediaSocialId: "1",
//		MerchantDescription:   "Gojek adalah cerdikiawan",
//	}
//	db.Create(&merchant)
//
}
