package repository

import(
	"github.com/davidnobels/loyalti-go-echo/src/database"
	"github.com/davidnobels/loyalti-go-echo/src/domain/model"
)

func GetOutlet() []model.Outlet {
	db := database.ConnectionDB()

	var outlet []model.Outlet
	db.Find(&outlet)
	db.Close()
	return outlet
}
