package main

import (
	"github.com/davidnobels/loyalti-go-echo/src/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/now"
	"time"
)


func main() {
	db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=loyalti.MerchantDb.Dev")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//connect postgre


	// Migrate the schema
	db.AutoMigrate(&model.SpecialProgram{})
	//db.AutoMigrate(&model.Outlet{},&model.Program{})
	var program = model.SpecialProgram{
		GeneralModels:      model.GeneralModels{
			Created:    time.Now(),
			CreatedBy:  "admin",
			Modified:   time.Now(),
			ModifiedBy: "admin",
			Active:     true,
			IsDeleted:  false,
			Deleted:    nil,
			Deleted_by: "",
		},
		ProgramName:        "Promo LAWSON",
		ProgramImage:       "",
		ProgramStartDate:   time.Now(),
		ProgramEndDate:     now.EndOfMonth(),
		ProgramDescription: "Diskon 50% setiap belanja 100rb",
		Card:               "Chop",
		OutletID:           2,
		MerchantId:         3,
	}

	//INSERT TO DATABASE, TABLE OUTLET
	//var outlet = model.Outlet{
	//	OutletName:"Starbucks FX Sudirman",
	//	OutletAddress:"FX Sudirman",
	//	OutletPhone:"(021) 30030377",
	//	OutletCity:"Kota Jakarta PUsat",
	//	OutletProvince:"DKI Jakarta",
	//	OutletPostalCode:"10270",
	//	OutletLongitude:"6.2250° S",
	//	OutletLatitude:"106.8038° E"}
	//
	//var outlet1 = model.Outlet{
	//	OutletName:"Starbucks FX Gading Serpong",
	//	OutletAddress:"Summarecon Mal Serpong",
	//	OutletPhone:"(021) 54200853",
	//	OutletCity:"Tangerang",
	//	OutletProvince:"Banten",
	//	OutletPostalCode:"15810",
	//	OutletLongitude:"106° 39' 56.1456'' E",
	//	OutletLatitude:"6° 19' 12.4968'' S"}
	//
	//var outlet2 = model.Outlet{
	//	OutletName:"Starbucks PP",
	//	OutletAddress:"Pacific Place",
	//	OutletPhone:"(021) 57973565",
	//	OutletCity:"Jakarta Selatan",
	//	OutletProvince:"DKI Jakarta",
	//	OutletPostalCode:"12190",
	//	OutletLongitude:"106.8100° E",
	//	OutletLatitude:"6.2247° S"}

	//merchant := model.Merchant{
	//	GeneralModels:        model.GeneralModels{
	//		Created:    time.Now(),
	//		CreatedBy:  "Admin",
	//		Modified:   time.Now(),
	//		ModifiedBy: "Admin",
	//		Active:     true,
	//		IsDeleted:  false,
	//		Deleted:    nil,
	//		Deleted_by: "",
	//	},
	//	MerchantName:         "PT MAP",
	//	MerchantEmail:        "MAP 30",
	//	MerchantPhoneNumber:  "08125125321",
	//	MerchantProvince:     "DKI Jakarta",
	//	MerchantCity:         "Jakarta",
	//	MerchantAddress:      "Sudirman Sahid Jaya kav 36 no 37",
	//	MerchantPostalCode:   "10120",
	//	MerchantCategory:     "Lifestyle",
	//	MerchantWebsite:      "www.map.co.id",
	//	MerchantMediaSocial:  "@mapindonesia",
	//	MerchantDescription:  "MAP is a lifestyle comany",
	//}
	//db.Create(&merchant)
	db.Create(&program)
	//var program = model.Program{
	//	GeneralModels:      model.GeneralModels{
	//		Created:    time.Now(),
	//		CreatedBy:  "admin",
	//		Modified:   time.Now(),
	//		ModifiedBy: "admin",
	//		Active:     true,
	//		IsDeleted:  false,
	//		Deleted:    nil,
	//		Deleted_by: "",
	//	},
	//	ProgramName:        "Gratis GoRIDE",
	//	ProgramImage:       "",
	//	ProgramStartDate:   time.Now(),
	//	ProgramEndDate:     time.Time{},
	//	ProgramDescription: "Gratis Go Ride selama bulan November",
	//	Card:               "Member",
	//	OutletID:           2,
	//	MerchantId:			2,
	//}
	////	program.MerchantId = 2
	//db.Create(&program)
	//db.Create(&outlet1)
	//db.Create(&outlet2)
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
