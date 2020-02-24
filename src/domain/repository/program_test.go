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

type SuiteProgram struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock

	repository RepositoryProgram
	program *model.Program
}

func (s *SuiteProgram) SetupSuiteProgram(){
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

	s.repository = CreateRepositoryProgram(s.DB)
	fmt.Println("test ini aman")
}

func (s *Suite) AfterTestProgram(_, _ string){
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInitProgram(t *testing.T){
	suite.Run(t, new(Suite))
}

type connection_mock_program struct {
	mock.Mock
}

//func (c connection_mock) ConnectionDB2() (gorm *gorm.DB){
//	fmt.Println("masuk mockingan connection")
//	return
//}

func (s *SuiteProgram) Test_Create_Program(){
	fmt.Println("test 1 aman")
	var (
			program = model.Program{
				Id: 1,
				ProgramName: "Diskon Mantap",
				ProgramDescription: "Beli 1 Gratis 1",
				Active:true,
			}
	)
	fmt.Println("test 2 aman")
	err := s.repository.CreateProgram(&program)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 aman", err)
}

func (s *SuiteProgram) Test_Update_Program(){
	fmt.Println("test update 1 aman")
	var (
		program = model.Program{
			Id:1,
			ProgramName:"Diskon Kurang Mantap",
		}
	)
	fmt.Println("test update 2 aman")
	err := s.repository.UpdateProgram(&program)
	fmt.Println("test 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 update aman")
}

func (s *SuiteProgram) Test_Delete_Program(){
	fmt.Println("test delete 1 aman")
	var (
		program = model.Program{
			Id: 1,
		}
	)
	fmt.Println("test delete 2 aman")
	err := s.repository.DeleteProgram(&program)
	fmt.Println("test delete 3 aman")
	require.NoError(s.T(), err)
	fmt.Println("test 4 delete aman")
}
