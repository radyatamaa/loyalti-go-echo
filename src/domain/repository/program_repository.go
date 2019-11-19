package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/davidnobels/loyalti-go-echo/src/database"
	"github.com/davidnobels/loyalti-go-echo/src/domain/model"
)

func GetProgram(page int, size int) []model.Program {
	db := database.ConnectionDB()

	var program []model.Program
	db.Find(&program)


	pagination.Paging(&pagination.Param{
		DB:	db,
		Page: page,
		Limit:	size,
		OrderBy:	[]string{"program_name desc"},
	}, &program)

	db.Close()
	return program
}

//func GetProgramByDate(page int, size int, sort int) []model.Program {
//	db := database.ConnectionDB()
//	sortBy := domain.Enum.OrderByDate_Asc
//
//	var program_date []model.Program
//	db.Find(&program_date)
//
//	pagination.Paging(&pagination.Param{
//		DB:      db,
//		Page:    page,
//		Limit:   size,
//		 //StartDate: start_date,
//		OrderBy: []string{" asc"} ,
//	}, &program_date)
//
//	db.Close()
//	return program_date
//}

func GetProgramByMerchantId(page int, size int) []model.Program {
	db := database.ConnectionDB()

	var program_by_merchant []model.Program

	db.Find(&program_by_merchant)

	 pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   size,
		OrderBy: []string{"merchant_id desc"} ,
	} ,&program_by_merchant)
	offset := ((page - 1) * (size))
	db.Table("merchants").Select("merchants.merchant_name, programs.program_name,programs.program_description").
		Joins("left join programs on programs.merchant_id = merchants.id").
		Order("merchant_id desc").Offset(offset).
		Scan(&program_by_merchant)

	test := append(program_by_merchant)

	fmt.Print(test)


	db.Close()
	return test
}