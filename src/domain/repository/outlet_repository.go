package repository

import(
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
)

func GetOutlet() []model.Outlet {
	db := database.ConnectionDB()

	var outlet []model.Outlet
	db.Find(&outlet)
	db.Close()
	return outlet
}
