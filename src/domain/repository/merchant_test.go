package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	merchant *model.NewMerchantCommand
}

func (s *Suite) SetupSuite(){
	fmt.Println("test 10 aman")
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
	fmt.Println("test 40 aman")

	s.DB.LogMode(true)
	fmt.Println("test 50 aman")

	s.repository = CreateRepository(s.DB)
	fmt.Println("test ini aman")
}

func (s *Suite) AfterTest(_, _ string){
	//require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T){
	suite.Run(t, new(Suite))
}

type connection_mock struct {
	mock.Mock
}

func (c connection_mock) ConnectionDB2() (gorm *gorm.DB){
	fmt.Println("masuk mockingan connection")
	return
}

func (s *Suite) Test_Create_Merchant(){
	fmt.Println("test 1 aman")
	var (
		merchant = model.NewMerchantTest{
			Id:            100,
			MerchantEmail: "abcd",
		}
	)
	fmt.Println("test 2 aman")
	//rows := sqlmock.Rows{}
	sql := regexp.QuoteMeta(`INSERT INTO "merchants" (id,merchant_email} VALUES (?,?)`)
	fmt.Println("sql : ",sql)
	a :=  s.mock.ExpectQuery(sql).WithArgs(merchant.Id,merchant.MerchantEmail).WillReturnRows(sqlmock.NewRows([]string{"merchant_email"}).AddRow(merchant.MerchantEmail))
	b := s.mock.ExpectExec(sql).WithArgs(merchant.Id, merchant.MerchantEmail).WillReturnError(errors.New("test error"))
	fmt.Println("isi a = ", a)
	fmt.Println("isi b = ", b)

	//s.mock.ExpectQuery(regexp.QuoteMeta(
	//	`INSERT INTO "merchants"("id","merchant_email")
	//		VALUES ($1, $2) RETURNING "merchants"."merchant_email"
	//		`,
	//	)).
	//	WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
	//	//WillReturnError(nil)
	//	WillReturnRows(
	//		//&rows,
	//		sqlmock.NewRows([]string{"merchant_email"}).
	//			AddRow(merchant.MerchantEmail),
	//		)
	fmt.Println("test 3 aman")
	err := s.repository.CreateMerchant(&merchant)
	fmt.Println("test 4 aman")

	if err == nil {
		fmt.Println("Unit Test Aman")
	}

	//require.NoError(s.T(), err)
	fmt.Println("test 5 aman")
}

//func Test_Create_Merchant (t *testing.T){
//	fmt.Println("masuk ke testing")
//	repo := Repo{}
//
//	merchant := model.NewMerchantCommand{}
//
//	//inisialisasi sqlmock
//	db, mock, err := sqlmock.New()
//	repo.DB = db
//
//	//mock agar insert tidak masuk ke DB
//	mock.ExpectQuery("INSERT INTO").WillReturnRows(sqlmock.NewRows([]string{"merchant_email"}))
//
//	err = repo.CreateMerchant(&merchant)
//
//	if err != nil {
//		//kasih tau error di repo.CreateMerchant
//		t.Errorf("Error : %s", err.Error())
//	}
//}