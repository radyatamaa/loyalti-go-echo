package main

import (
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	//"time"
)
//type User struct {
//	Id int `json:"id","AUTO_INCREMENT"`
//	Name string `json:"name"`
//}

func main() {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	//connect postgre

	// Migrate the schema
	db.AutoMigrate(&model.MerchantCategory{}, &model.Merchant{})
	//db.AutoMigrate(&User{})

	//category := model.MerchantCategory{
	//	GeneralModels: model.GeneralModels{
	//		Created:    time.Now(),
	//		CreatedBy:  "Admin",
	//		Modified:   time.Now(),
	//		ModifiedBy: "Admin",
	//		Active:     false,
	//		IsDeleted:  false,
	//		Deleted:    nil,
	//		Deleted_by: "",
	//	},
	//	CategoryName:  "Lifestyle",
	//	ImageUrl:      "",
	//}
	//db.Create(&category)

	//var merchant1 = model.Merchant{
	//	GeneralModels: model.GeneralModels{
	//		Created:    time.Now(),
	//		CreatedBy:  "Admin",
	//		Modified:   time.Now(),
	//		ModifiedBy: "Admin",
	//		Active:     true,
	//		IsDeleted:  false,
	//		Deleted:    nil,
	//		Deleted_by: "",
	//	},
	//	MerchantName: "Bateeq",
	//	MerchantPhoneNumber: "+62 821 373 92222",
	//	MerchantEmail:"contact@bateeq.com",
	//	MerchantWebsite:"www.bateeq.com",
	//	MerchantMediaSocial:"@bateeqshop",
	//	MerchantAddress:"Equity Tower, floor 15th, Suite C SCBD Lot 9. Jalan Jend. Sudirman kav 52-53",
	//	MerchantCity:"Jakarta Selatan",
	//	MerchantPostalCode:"12190",
	//	MerchantProvince:  "DKI Jakarta",
	//	MerchantDescription: " bateeq is an Indonesian fashion brand that offers a fresh, fashion-forward take on batik through our clothing line for men, women and children",

	//}
	//	merchant1.MerchantCategoryId = 1
	//	db.Create(&merchant1)

}
