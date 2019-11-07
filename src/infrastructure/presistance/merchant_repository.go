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
	db.Close()
	return merchant
}

func CreateMerchant(m domain.Merchant)uint {
	db := database.ConnectionDB()
	merchant := m

	db.Create(merchant)
	db.Close()
	return merchant.ID

}

