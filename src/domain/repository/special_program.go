package repository

import (
	"fmt"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func GetSpecialProgram(page *int, size *int, sort *int) []model.SpecialProgram {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var special []model.SpecialProgram
	db.Find(&special)

	if page != nil && size != nil {
		pagination.Paging(&pagination.Param{
			DB:      db,
			Page:    *page,
			Limit:   *size,
			OrderBy: []string{"program_name desc"},
		}, &special)

		//var enum domain.List
		if sort != nil {
			fmt.Println("masuk gak kesini")
			switch *sort {
			case 1 :
				db.Order("created asc").Find(&special)
			case 2:
				db.Order("created desc").Find(&special)
			case 3:
				db.Order("program_name asc").Find(&special)
			case 4:
				db.Order("program_name desc").Find(&special)
			}


			//switch *sort {
			//case enum.OrderByDate_Asc:
			//	db.Order("created asc").Find(&program)
			//	fmt.Println("test asc")
			//case enum.OrderByDate_Desc:
			//	db.Order("created desc").Find(&program)
			//	}
		}
	}
	db.Close()
	return special
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

//func GetProgramByMerchantId(page int, size int) []model.Program {
//	db := database.ConnectionDB()
//
//	var program_by_merchant []model.Program
//
//	db.Find(&program_by_merchant)
//
//	 pagination.Paging(&pagination.Param{
//		DB:      db,
//		Page:    page,
//		Limit:   size,
//		OrderBy: []string{"merchant_id desc"} ,
//	} ,&program_by_merchant)
//	offset := ((page - 1) * (size))
//	db.Table("merchants").Select("merchants.merchant_name, programs.program_name,programs.program_description").
//		Joins("left join programs on programs.merchant_id = merchants.id").
//		Order("merchant_id desc").Offset(offset).
//		Scan(&program_by_merchant)
//
//	test := append(program_by_merchant)
//
//	fmt.Print(test)
//
//
//	db.Close()
//	return test
//}