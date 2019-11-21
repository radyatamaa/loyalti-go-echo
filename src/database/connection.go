package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectionDB() *gorm.DB{
	db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=loyalti.MerchantDb.Dev")
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	return db
}

//func ConnectPostgre() *gorm.DB {
//	db, err := gorm.Open("postgres", "host=localhost port=5432 user=sa dbname=loyalti-postgre sslmode=disable password=Moonlay2019.")
//	//db, err := gorm.Open("postgres", "postgres:postgres@/loyalti")
//	//defer db.Close()
//
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	return db
//}