package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
)

func GetSocialMedia(page *int, size *int) []model.MerchantSocialMedia {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var sosmed []model.MerchantSocialMedia
	db.Find(&sosmed)

	if size != nil && page != nil {
	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: *page,
		Limit:	*size,
		OrderBy:	[]string{"category_name desc"},
	}, &sosmed)
	}
	db.Close()
	return sosmed
}
