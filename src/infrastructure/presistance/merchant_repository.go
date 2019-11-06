package presistance

import (
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/domain"
)

func GetAll() []domain.Merchant {
	db, err := gorm.Open("mssql", "sqlserver://sa:Moonlay2019.@11.11.5.146?database=loyalti.MerchantDb.Dev")
	if err != nil {
		panic("failed to connect database")

		//connect to localhost
		/*db, err := gorm.Open("mssql", "sqlserver://@mssqllocaldb?database=loyalti.MerchantDb.Dev")
		if err != nil {
			panic("failed to connect database")*/
	}
	//if err != nil {
	//	return nil, err
	//}
	defer db.Close()
	var merchant []domain.Merchant
	db.Find(&merchant)

	return merchant
}
