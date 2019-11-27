package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetProgram(page *int, size *int, sort *int, category *int) []model.Program {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var program []model.Program
	db.Find(&program)

	if sort !=nil {
		switch *sort {
		case 1:
			if page != nil && size != nil {
				pagination.Paging(&pagination.Param{
					DB: 	db,
					Page: *page,
					Limit: *size,
					OrderBy: []string{"created asc"},
				}, &program)
				db.Order("created asc").Where("category_id=?",category).Find(&program)
			}
		case 2:
			if page != nil && size != nil {
				pagination.Paging(&pagination.Param{
					DB: 	db,
					Page: *page,
					Limit: *size,
					OrderBy: []string{"created desc"},
				}, &program)
			}
		case 3:
			if page != nil && size != nil {
				pagination.Paging(&pagination.Param{
					DB: 	db,
					Page: *page,
					Limit: *size,
					OrderBy: []string{"program_name asc"},
				}, &program)
			}
		case 4:
			if page != nil && size != nil {
				pagination.Paging(&pagination.Param{
					DB: 	db,
					Page: *page,
					Limit: *size,
					OrderBy: []string{"program_name desc"},
				}, &program)
			}
		}
	}
	db.Close()
	return program
}