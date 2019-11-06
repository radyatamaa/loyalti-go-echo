package database

import "github.com/jinzhu/gorm"

func ConnectionDB() *gorm.DB{
	db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=loyalti.MerchantDb.Dev")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	return db
}
