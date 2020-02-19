package repository

import (
	"fmt"
	"github.com/beevik/guid"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

type EmployeeRepository interface {
	CreateEmployee (newemployee *model.Employee) error
}

type employee_repo struct {
	DB *gorm.DB
	database.Connection_Interface
}

func (p employee_repo) CreateEmployee (newemployee *model.Employee) error {
	fmt.Println("masuk fungsi")
	employee := model.Employee{
	}
	fmt.Println("error nil")

	err := p.DB.Create(&employee).Error
	if err != nil {
		fmt.Println("Error DB.Create", err.Error())
	}
	fmt.Println("selamat/sukses")

	return err
}

func CreateEmployeeRepository (db *gorm.DB) EmployeeRepository {
	return &employee_repo{
		DB:db,
	}
}

func CreateEmployee(employee *model.Employee) string{
	db := database.ConnectionDB()
	employee.Id = guid.NewString()
	db.Create(&employee)
	return employee.EmployeeEmail
}


func UpdateEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	db.Model(&employee).Where("employee_email = ?", employee.EmployeeEmail).Update(&employee)
	return employee.EmployeeEmail
}


func DeleteEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	db.Model(&employee).Where("outlet_id = ?", employee.OutletId).Delete(&employee)
	//db.Model(&employee).Where("employee_email = ?",employee.EmployeeEmail).Update("active", false)
	return "berhasil dihapus"
}

func GetEmployee(page *int, size *int, sort *int) []model.Employee {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var employee []model.Employee
	db.Find(&employee)



	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"employee_name desc"},
				}, &employee)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"employee_name asc"},
				}, &employee)
			}
		}
	}
	db.Close()
	return employee
}
