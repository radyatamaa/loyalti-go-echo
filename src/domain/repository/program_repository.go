package repository

import (
	"database/sql"
	"fmt"

	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/sirupsen/logrus"
)

func CreateProgram(program *model.Program) string {
	db := database.ConnectionDB()

	programObj := *program

	db.Create(&programObj)

	return programObj.ProgramName
}

func UpdateProgram(program *model.Program) string {
	db := database.ConnectionDB()
	db.Model(&program).Where("id = ?", program.Id).Update(&program)
	return "Berhasil diUpdate"
}

func DeleteProgram(program *model.Program) string {
	db := database.ConnectionDB()
	db.Model(&program).Where("id= ?", program.Id).Delete(&program)
	return "berhasil dihapus"
}

func TotalPoint(id *int, pay *int) int {
	db := database.ConnectionDB()
	program := model.Program{}
	db.Model(&program).Where("id = ?", id).Find(&program)
	if *pay < *(program.MinPayment) {
		fmt.Println("Customer Tidak mendapatkan Point")
		return 0
	}
	var total = *pay * *(program.ProgramPoint) / *(program.MinPayment)
	fmt.Println("Total Point : ", total)
	return total
}

func GetProgram(page *int, size *int, sort *int, category *int, id *int) []model.Program {

	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var program []model.Program
	var rows *sql.Rows
	var err error
	var total int

	if sort != nil {
		switch *sort {
		case 1:
			if page != nil && size != nil && category == nil {
				rows, err = db.Find(&program).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil {
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		case 2:
			if page != nil && size != nil && category == nil {
				rows, err = db.Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil {
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		case 3:
			if page != nil && size != nil && category == nil {
				rows, err = db.Find(&program).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil {
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		case 4:
			if page != nil && size != nil && category == nil {
				rows, err = db.Find(&program).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil {
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		}
	} else {
		if page != nil && size != nil {
			rows, err = db.Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		} else {
			rows, err = db.Find(&program).Rows()
			if err != nil {
				panic(err)
			}
		}
	}
	if id != nil {
		rows, err = db.Where("id = ?", id).First(&program).Rows()
		if err != nil {
			panic(err)
		}
	}

	result := make([]model.Program, 0)

	for rows.Next() {
		t := &model.Program{}
		fmt.Println(t)
		benefitmemory := t.Benefit
		fmt.Println(&benefitmemory)

		err = rows.Scan(
			&t.Id,
			&t.Created,
			&t.CreatedBy,
			&t.Modified,
			&t.ModifiedBy,
			&t.Active,
			&t.IsDeleted,
			&t.Deleted,
			&t.Deleted_by,
			&t.ProgramName,
			&t.ProgramImage,
			&t.ProgramStartDate,
			&t.ProgramEndDate,
			&t.ProgramDescription,
			&t.Card,
			&t.OutletID,
			&t.MerchantId,
			&t.CategoryId,
			&t.Benefit,
			&t.TermsAndCondition,
			&t.Tier,
			&t.RedeemRules,
			&t.RewardTarget,
			&t.QRCodeId,
			&t.ProgramPoint,
			&t.MinPayment,
		)
		merchant := new(model.Merchant)

		db.Table("merchants").
			Select("merchants.merchant_name").
			Where("id = ?", t.MerchantId).
			First(&merchant)
		t.MerchantName = merchant.MerchantName
		if err != nil {
			logrus.Error(err)
			return nil
		}
		result = append(result, *t)
	}
	db.Close()
	return result
}
