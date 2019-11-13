package main

import (
	"github.com/davidnobels/loyalti-go-echo/src/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
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
	db.AutoMigrate(&domain.Outlet{}, &domain.Program{})


	//INSERT TO DATABASE, TABLE OUTLET
	var outlet = domain.Outlet{ OutletName:"Starbucks FX Sudirman",
								OutletAddress:"FX Sudirman",
								OutletPhone:"(021) 30030377",
								OutletCity:"Kota Jakarta PUsat",
								OutletProvince:"DKI Jakarta",
								OutletPostalCode:"10270",
								OutletLongitude:"6.2250° S",
								OutletLatitude:"106.8038° E"}

	db.Create(&outlet)
	//var program = domain.Program{ ProgramName:"Belanja 11:11", ProgramDescription:"Program baru", Card:"Chop"}
	//db.Create(&program)
	//Create Insert to database
	//var card = domain.CardType{ CardTypeName:"Chop"}
	//db.Create(&card)

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
