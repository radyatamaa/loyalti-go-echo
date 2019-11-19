package repository

import (

)

func GetCard(page int, size int) []model.CardType {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var card []model.CardType
	db.Find(&card)
	db.Close()

	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: page,
		Limit:	size,
		OrderBy:	[]string{"program_name desc"},
	}, &card)

	return card
}