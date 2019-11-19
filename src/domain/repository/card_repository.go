package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetCard(page *int, size *int) []model.CardType {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var card []model.CardType
	db.Find(&card)

	if size != nil && page != nil {
	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: *page,
		Limit:	*size,
		OrderBy:	[]string{"program_name desc"},
		}, &card)
	}
	db.Close()
	return card
}