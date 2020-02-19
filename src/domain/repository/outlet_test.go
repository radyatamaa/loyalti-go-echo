package repository

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OutletSuite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	outlet_repository OutletRepository
	outlet *model.Outlet
}

func (s *OutletSuite) SetupSuite(){
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

	s.outlet_repository = CreateOutletRepository(s.DB)
	fmt.Println("test ini aman")
}

func (s *OutletSuite) AfterTest(_, _ string){

}

func TestInitOutlet(t *testing.T){
	suite.Run(t, new(OutletSuite))
}

type outlet_connection_mock struct {
	mock.Mock
}

func (c outlet_connection_mock) ConnectionDB2() (gorm *gorm.DB){
	fmt.Println("masuk connection mock")
	return
}

func (s *OutletSuite) Test_Create_Outlet(){
	fmt.Println("test 1 aman")
	var (
		outlet = model.Outlet{
			OutletName:       "Tokopedia",
			OutletAddress:    "JKt",
		}
	)
	fmt.Println("test 2 aman")

	fmt.Println("lewat yg pertama")
	fmt.Println("lewat yg kedua")

	fmt.Println("test 3 aman")
	err := s.outlet_repository.CreateOutlet(&outlet)
	fmt.Println("test 4 aman")

	if err == nil {
		fmt.Println("Unit test Aman")
	}

	fmt.Println("test 5 aman")
}
