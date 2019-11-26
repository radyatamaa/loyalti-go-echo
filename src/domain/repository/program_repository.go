package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetProgram(page *int, size *int, sort *int) []model.Program {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var program []model.Program
	db.Find(&program)

	if page != nil && size != nil {
		pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    *page,
			Limit:   *size,
			OrderBy: []string{"program_name desc"},
		}, &program)

		if sort != nil {
			fmt.Println("masuk gak kesini")
			switch *sort {
			case 1:
				db.Order("created asc").Find(&program)
				if page != nil && size != nil {
					pagination.Paging(&pagination.Param{
						DB:      db,
						Page:    *page,
						Limit:   *size,
						OrderBy: []string{"created asc"},
					}, &program)
				}
			case 2:
				db.Order("created desc").Find(&program)
			case 3:
				db.Order("program_name asc").Find(&program)
			case 4:
				db.Order("program_name desc").Find(&program)
			case 5:
				db.Find(&program)
			}
		}
	}
	db.Close()
	return program
}