package repository

import (
	"database/sql"
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/sirupsen/logrus"
)


func CreateProgram(program *model.Program) string{
	db := database.ConnectionDB()
	programObj := *program
	db.Create(&programObj)
	return programObj.ProgramName
}

func UpdateProgram  (program *model.Program) string {
	db := database.ConnectionDB()
	db.Model(&program).Where("id = ?", program.Id).Update(&program)
	return "Berhasil diUpdate"
}

func DeleteProgram(program *model.Program) string {
	db := database.ConnectionDB()
	db.Model(&program).Where("id= ?", program.Id).Update("active", false)
	return "berhasil dihapus"
}

//func FindEmployee (outletid string) []model.TotalPoint{
//	db := database.ConnectionDB()
//	employee := model.Employee{}
//	db.Model(&employee).Where("outlet_id = ?", outletid).Find(&employee)
//	return nil
//}

func TotalPoint(id int, pay int, pin string, outletid string, cardtype string) []model.TotalPoint {
	db := database.ConnectionDB()
	employee := model.Employee{}
	totalpoint := []model.TotalPoint{}
	program := model.Program{}
	db.Model(&program).Where("id = ?", id).Find(&program)
	db.Model(&employee).Where("employee_pin = ?", pin).Find(&employee)
	db.Model(&program).Where("outlet_id = ?", outletid).Find(&program)
	db.Model(&program).Where("card = ?", cardtype).Find(&program)

	if cardtype != program.Card {
		fmt.Println("Customer tidak mendapatkan point dengan kartu ini")
		return nil
	}

	if outletid != program.OutletID  {
		fmt.Println("Anda tidak memiliki akses ")
		return nil
	}
	if  pin != employee.EmployeePin {
		fmt.Println("Pin Anda Salah. Silahkan Coba Lagi ")
		return nil
	}
	if outletid == employee.OutletId && pin == employee.EmployeePin {
		if pay < *(program.MinPayment)  {
			fmt.Println("Customer tidak mendapatkan poin ")
			return nil
		}
		var total = pay * *(program.ProgramPoint) / *(program.MinPayment)
		t := &model.TotalPoint{}
		t.Total = total
		updatepoint := append(totalpoint, *t)
		fmt.Printf("Customer mendapatkan %d poin \n", total)
		return updatepoint
	}
	return nil
}

func TotalChop(id int, pay int, pin string, outletid string, cardtype string) []model.TotalChop {
	db := database.ConnectionDB()
	employee := model.Employee{}
	totalchop := []model.TotalChop{}
	program := model.Program{}
	db.Model(&program).Where("id = ?", id).Find(&program)
	db.Model(&employee).Where("employee_pin = ?", pin).Find(&employee)
	db.Model(&program).Where("outlet_id = ?", outletid).Find(&program)
	db.Model(&program).Where("card = ?", cardtype).Find(&program)

	if cardtype != program.Card {
		fmt.Println("Customer tidak mendapatkan chop dengan kartu ini")
		return nil
	}

	if outletid != program.OutletID  {
		fmt.Println("Anda tidak memiliki akses ")
		return nil
	}
	if  pin != employee.EmployeePin {
		fmt.Println("Pin Anda Salah. Silahkan Coba Lagi ")
		return nil
	}
	if outletid == employee.OutletId && pin == employee.EmployeePin {
		if pay < *(program.MinPayment)  {
			fmt.Println("Customer tidak mendapatkan tambahan chop ")
			return nil
		}

		if pay >= *(program.MinPayment){
		var total = pay / pay * 1
		t := &model.TotalChop{}
		t.Total = total
		updatechop := append(totalchop, *t)
		fmt.Printf("Customer mendapatkan %d chop \n", total)
		return updatechop
		}
	}
	return nil
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
			if page != nil && size != nil && category == nil{
				rows, err = db.Find(&program).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil{
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		case 2:
			if page != nil && size != nil && category == nil{
				rows, err = db.Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil{
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		case 3:
			if page != nil && size != nil && category == nil{
				rows, err = db.Find(&program).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil{
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("program_name asc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		case 4:
			if page != nil && size != nil && category == nil{
				rows, err = db.Find(&program).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
			if category != nil && page != nil && size != nil{
				rows, err = db.Where("category_id = ?", category).Find(&program).Order("program_name desc").Count(total).Limit(*size).Offset(*page).Rows()
				if err != nil {
					panic(err)
				}
			}
		}
	}else {
		if page != nil && size != nil {
			rows, err = db.Find(&program).Order("created desc").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		} else{
			rows, err = db.Find(&program).Rows()
			if err != nil {
				panic(err)
			}
		}
	}
	if id != nil {
		rows, err = db.Where("id = ?", id).First(&program).Rows()
		if err != nil{
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
		//t.MerchantName = merchant.MerchantName
		if err != nil {
			logrus.Error(err)
			return nil
		}
		result = append(result,*t)
		}
		db.Close()
	return result
	}
