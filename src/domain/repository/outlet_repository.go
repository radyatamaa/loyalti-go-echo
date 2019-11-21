package repository

import(
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetOutlet(page *int, size *int) []model.Outlet {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var outlet []model.Outlet
	db.Find(&outlet)

	if page != nil && size != nil {
	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: *page,
		Limit:	*size,
		OrderBy:	[]string{"outlet_name desc"},
	}, &outlet)
	}
	db.Close()
	return outlet
}
