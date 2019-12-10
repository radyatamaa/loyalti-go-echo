package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

func CreateEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	employeeObj := *employee
	db.Create(&employeeObj)
	return employeeObj.EmployeeEmail
}

func UpdateEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	db.Model(&employee).Where("employee_email = ?", employee.EmployeeEmail).Update(&employee)
	return employee.EmployeeEmail
}

func DeleteEmployee(employee *model.Employee) string {
	db := database.ConnectionDB()
	db.Model(&employee).Where("employee_email = ?",employee.EmployeeEmail).Update("active", false)
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
