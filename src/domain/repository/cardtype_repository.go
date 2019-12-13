package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func CreateCard(card *model.CardType) string{
	db := database.ConnectionDB()
	cardObj := *card
	db.Create(&card)
	return cardObj.CardTypeName
}

func UpdateCard(card *model.CardType) string {
	db := database.ConnectionDB()
	db.Model(&card).Where("id = ?", card.Id).Update(&card)
	return card.CardTypeName
}

func DeleteCard(card *model.CardType) string {
	db := database.ConnectionDB()
	db.Model(&card).Where("id = ?", card.Id).Update("active", false)
	return "berhasil dihapus"
}

func GetCard(page *int, size *int, sort *int) []model.CardType {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var card []model.CardType
	db.Find(&card)

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"card_type_name desc"},
				}, &card)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"card_type_name asc"},
				}, &card)
			}
		}
	}
	db.Close()
	return card
}
