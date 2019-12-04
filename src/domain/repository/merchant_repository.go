package repository

import (
	"github.com/biezhi/gorm-paginator/pagination"
	//"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	//"net/http"
)

func CreateMerchant(merchant *model.Merchant) string{
	db := database.ConnectionDB()

	merchantObj := *merchant

	db.Create(&merchantObj)

	return merchantObj.MerchantEmail

}

func UpdateMerchant(merchant *model.Merchant) func(echo.Context) error {
	db := database.ConnectionDB()
	return func(c echo.Context) error {
		name := c.Param("merchant_name")
		email := c.Param("merchant_email")
		phone := c.Param("merchant_phone_number")
		category := c.Param("merchant_category_id")
		website := c.Param("merchant_website")
		address := c.Param("merchant_address")
		city := c.Param("merchant_city")
		postalcode := c.Param("merchant_postal_code")
		province := c.Param("merchant_province")
		socialmedia := c.Param("merchant_social_media_id")
		description := c.Param("merchant_description")
		imageprofile := c.Param("merchant_image_profile")
		gallery := c.Param("merchant_gallery")

		db.Where("merchant_email = ? ",email).Find(&merchant)

		merchant.MerchantName = name
		merchant.MerchantEmail = email
		merchant.MerchantPhoneNumber = phone
		merchant.MerchantCategoryId = category
		merchant.MerchantWebsite = website
		merchant.MerchantAddress = address
		merchant.MerchantCity = city
		merchant.MerchantPostalCode = postalcode
		merchant.MerchantProvince = province
		merchant.MerchantMediaSocialId = socialmedia
		merchant.MerchantDescription = description
		merchant.MerchantImageProfile = imageprofile
		merchant.MerchantGallery = gallery

		db.Save(&merchant)
		return nil

	}
}

func DeleteMerchant(merchant *model.Merchant, email string) string {
	db := database.ConnectionDB()

	merchantObj := *merchant

	db.Where("merchant_email = ? ", email).Delete(&merchantObj)

	return merchantObj.MerchantEmail

}


func GetMerchant(page *int, size *int, sort *int, email *string) []model.Merchant {
	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var merchant []model.Merchant
	db.Find(&merchant)

	if sort != nil {
		switch *sort {
		case 1:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name desc"},
				}, &merchant)
			}
		case 2:
			if size != nil && page != nil {
				pagination.Paging(&pagination.Param{
					DB:      db,
					Page:    *page,
					Limit:   *size,
					OrderBy: []string{"merchant_name asc"},
				}, &merchant)
			}
		}
	}

	if email != nil {
		if page != nil && size != nil {
			pagination.Paging(&pagination.Param{
				DB: 	db,
				Page: *page,
				Limit: *size,
				OrderBy: []string{"merchant_name asc"},

			}, &merchant)

			db.Order("id").Where("merchant_email = ?", email).Find(&merchant)
			//fmt.Println("2")
			//rows, err = db.Where("merchant_id = ?", id).Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			//if err != nil {
			//	panic(err)
			//}
		} else{
			db.Order("id").Where("merchant_email = ?", email).Find(&merchant)
		}
	}



	db.Close()
	return merchant
}
