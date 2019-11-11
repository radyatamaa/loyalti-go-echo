package repository

import(
	"github.com/jinzhu/gorm"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
)

func GetProgram() []model.Program {
	db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=loyalti.MerchantDb.Dev")
	if err != nil {
		panic("failed to connect database")
	}

	var program []model.Program
	db.Find(&program)

	return program
}
