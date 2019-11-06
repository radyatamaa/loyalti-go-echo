package presistance

import (
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain"
)

func GetAll() []domain.Merchant {
	db := database.ConnectionDB()
	//if err != nil {
	//	return nil, err
	//}
	var merchant []domain.Merchant
	db.Find(&merchant)

	return merchant
}

