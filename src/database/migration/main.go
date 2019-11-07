package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/davidnobels/loyalti-go-echo/src/domain"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
	db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=loyalti.MerchantDb.Dev")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//connect postgre


	// Migrate the schema
	db.AutoMigrate(&domain.Merchant{})

	//mengurangi kolom
	//db.Model(&domain.Merchant{}).DropColumn("row_status")
	//db.Model(&domain.Merchant{}).DropColumn("created_at")
	//db.Model(&domain.Merchant{}).DropColumn("created_by")
	//db.Model(&domain.Merchant{}).DropColumn("updated_by")
	//db.Model(&domain.Merchant{}).DropColumn("updated_at")
	//db.Model(&domain.Merchant{}).DropColumn("deleted_by")
	//db.Model(&domain.Merchant{}).DropColumn("deleted_at")

	//if err == nil{
	//	panic("success migration")
	//}

	//var merchant []Merchant
	//db.Find(&merchant)
	//fmt.Println(merchant)
	//// Create
	//db.Create(&Product{Code: "L1212", Price: 1000})
	//
	//// Read
	//var product Product
	//db.First(&product, 1) // find product with id 1
	//db.First(&product, "code = ?", "L1212") // find product with code l1212
	//
	//// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// Delete - delete product
	//db.Delete(&product)
}
