package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/beevik/guid"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EmployeeSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	employee_repository EmployeeRepository
	employee *model.Employee
}

func (s *EmployeeSuite) SetupSuite(){
	fmt.Println("test suite aman")
	var (
		db *sql.DB
		err error
	)
	fmt.Println("test 20 aman")
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	fmt.Println("test 30 aman")

	s.DB, err = gorm.Open("sqlserver", db)
	require.NoError(s.T(), err)
	fmt.Println("test40 aman")

	s.DB.LogMode(true)
	fmt.Println("test50 aman")

	s.employee_repository = CreateEmployeeRepository(s.DB)
	fmt.Println("test ini aman")
}

func (s *EmployeeSuite) AfterTest(_, _ string){

}

func TestInitEmployee(t *testing.T){
	suite.Run(t, new(EmployeeSuite))
}

type employee_connection_mock struct {
	mock.Mock
}

func (c employee_connection_mock) ConnectionDB() (gorm *gorm.DB) {
	fmt.Println("masuk connection mock")
	return
}

func (s *EmployeeSuite) Test_Create_Employee(){
	fmt.Println("test 1 aman")
	var (
		employee = model.Employee{
			Id:            guid.NewString(),
			Created:       time.Now(),
			CreatedBy:     "Admin",
			Modified:      time.Now(),
			ModifiedBy:    "Admin",
			Active:        true,
			IsDeleted:     false,
			Deleted:       nil,
			Deleted_by:    "",
			EmployeeName:  "David",
			EmployeeEmail: "david@email.com",
			EmployeePin:   "969696",
			EmployeeRole:  "Cashier",
			OutletId:      "1",
			OutletName:    "Bu Nanik",
		}
	)
	fmt.Println("test 2 aman")
	err := s.employee_repository.CreateEmployee(&employee)
	fmt.Println("test 3 aman")

	if err == nil {
		fmt.Println("Unit Test Aman")
	}

	fmt.Println("test 4 aman")
}