package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/felixsiburian/loyalti-go-echo/src/database"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
)

func GetProgram(page int, size int) []model.Program {
	db := database.ConnectionDB()

	var program []model.Program
	db.Find(&program)
	db.Close()

	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: page,
		Limit:	size,
		OrderBy:	[]string{"program_name desc"},
	}, &program)

	return program
}