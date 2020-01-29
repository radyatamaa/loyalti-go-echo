package main

import (
	"github.com/beevik/guid"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"

	//"github.com/radyatamaa/loyalti-go-echo/src/database"

)

func main() {
	db := database.ConnectionDB()
	employee := model.Employee{
		Id:            guid.NewString(),
		Created:       time.Now(),
		CreatedBy:     "",
		Modified:      time.Now(),
		ModifiedBy:    "",
		Active:        true,
		IsDeleted:     false,
		Deleted:       nil,
		Deleted_by:    "",
		EmployeeName:  "Patro Cashback",
		EmployeeEmail: "PatroCashback@gmail.com",
		EmployeePin:   "123123",
		EmployeeRole:  "Manager",
		OutletId:      "2",
		OutletName:    "Chatime fx",
	}
	db.Create(&employee)
	//db.AutoMigrate(&employee)
	}



