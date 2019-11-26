package repository

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/echo.v2"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"net/http"
)

func newMerchant(db *gorm.DB) func(echo.Context) error {
	return func (c echo.Context) error {
		merchant_name := c.Param("merchant_name")
		merchant_phone_number := c.Param("merchant_phone_number")
		merchant_email := c.Param("merchant_email")

		db.Create(&model.Merchant{
			MerchantName: merchant_name,
			MerchantPhoneNumber: merchant_phone_number,
			MerchantEmail: merchant_email,
		})
		return c.String(http.StatusOK, merchant_name+"user created succesfully")
	}
}
