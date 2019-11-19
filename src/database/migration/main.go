package main

import (
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"time"
)
//type User struct {
//	Id int `json:"id","AUTO_INCREMENT"`
//	Name string `json:"name"`
//}

func main() {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	//connect postgre


	var merchant1 = model.Merchant{
		GeneralModels: model.GeneralModels{
			Created:    time.Now(),
			CreatedBy:  "Admin",
			Modified:   time.Now(),
			ModifiedBy: "Admin",
			Active:     true,
			IsDeleted:  false,
			Deleted:    nil,
			Deleted_by: "",
		},
		MerchantName: "GoJek",
		MerchantPhoneNumber: "+62 821 373 92222",
		MerchantEmail:"contact@gojek.com",
		MerchantWebsite:"www.gojek.com",
		MerchantAddress:"Gojek Tower",
		MerchantCity:"Jakarta ",
		MerchantPostalCode:"11190",
		MerchantProvince:  "DKI Jakarta",
		MerchantDescription: "Jo adalah cerdikiawan",
		MerchantMediaSocialId: 1,
		MerchantCategoryId: 6,
	}
		db.Create(&merchant1)
}
