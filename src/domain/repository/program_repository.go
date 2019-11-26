package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetProgram(page *int, size *int,sort *int, category *int) []model.Program {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var program []model.Program
	db.Find(&program)

	if sort != nil{
		switch *sort {
		case 1 :
			if size != nil && page != nil && category !=nil  {
				fmt.Println("masuk ke if")
				pagination.Paging(&pagination.Param{
					DB:	db,
					Page: *page,
					Limit:	*size,
					OrderBy:	[]string{"program_name asc"},
				}, &program)
				fmt.Println("harusnya order terbaca")
				db.Order("program_name asc").Where("category_id = ? ", category).Find(&program)
			}
		case 2:
			if size != nil && page != nil && category !=nil  {
				fmt.Println("masuk dong")
				pagination.Paging(&pagination.Param{
					DB:	db,
					Page: *page,
					Limit:	*size,
					OrderBy:	[]string{"program_name desc"},
				}, &program)
				db.Order("program_name desc").Where("category_id = ? ", category).Find(&program)
			}
		case 3:
			if size != nil && page != nil && category !=nil  {
				fmt.Println("masuk dong")
				pagination.Paging(&pagination.Param{
					DB:	db,
					Page: *page,
					Limit:	*size,
					OrderBy:	[]string{"program_name desc"},
				}, &program)
				db.Order("program_start_date asc").Where("category_id = ? ", category).Find(&program)
			}
		case 4:
			if size != nil && page != nil && category !=nil  {
				fmt.Println("masuk dong")
				pagination.Paging(&pagination.Param{
					DB:	db,
					Page: *page,
					Limit:	*size,
					OrderBy:	[]string{"program_name desc"},
				}, &program)
				db.Order("program_start_date desc").Where("category_id = ? ", category).Find(&program)
			}
		}
	}

	db.Close()
	return program
}
