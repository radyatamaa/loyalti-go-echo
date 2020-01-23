package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/stretchr/testify/mock"
	"testing"
)

type Connection_Mock struct {
	mock.Mock
}

func (c *Connection_Mock) ConnectionDB2() *gorm.DB {
	fmt.Println("masuk ke mockingan")
	db := gorm.DB{}
	return &db
}

func (c *Connection_Mock) Create(value interface{}) *gorm.DB{
	fmt.Println("masuk ke mocking create")
	//db := gorm.DB{}
	return nil
}


func TestCreateMerchant(t *testing.T) {
	test := new(Connection_Mock)
	merchant := Merchant{test}
	merch := model.Merchant{}
	//db := gorm.DB{}
	test.On("Create", &merch).Return(nil)
	str := merchant.CreateMerchant2(&merch)
	fmt.Println(str)
	test.AssertExpectations(t)
}
