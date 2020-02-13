package repository

import (
	"database/sql"
	"fmt"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func CreateOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	outlet.Created = time.Now()
	outlet.Modified = time.Now()
	outletObj := *outlet

	db.Create(&outletObj)

	return outletObj.OutletName
}

func UpdateOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	db.Model(&outlet).Where("id = ?", outlet.Id).Update(&outlet)
	return outlet.OutletName
}

func DeleteOutlet(outlet *model.Outlet) string {
	db := database.ConnectionDB()
	db.Model(&outlet).Where("id= ?", outlet.Id).Update("active", false)
	return "berhasil dihapus"
}

func GetOutlet(page *int, size *int, id *int, email *string) []model.Outlet {

	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var outlet []model.Outlet
	var rows *sql.Rows
	var err error
	var total int
	var merchantID int
	merchant_email := model.Merchant{}
	db.Model(&merchant_email).Where("merchant_email = ?", email).Find(&merchant_email)
	if email != nil {
		merchantID = merchant_email.Id
	}
	fmt.Println("email", email)

	if id != nil {

		if page != nil && size != nil  {
			fmt.Println("2")
			rows, err = db.Where("merchant_id = ?", id).Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
			fmt.Println("Error : ", err.Error())
      }
		}
	} else if  merchantID != 0 {

		if page != nil && size != nil {
			//fmt.Println("2")
			rows, err = db.Where("merchant_id = ?", merchantID).Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				fmt.Println("Error : ", err.Error())
			}
		}
	} else {
		if page != nil && size != nil {
			rows, err = db.Find(&outlet).Order("outlet_name").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}
		} else {
			rows, err = db.Find(&outlet).Rows()
			if err != nil {
				panic(err)
			}
		}
	}
	if id != nil {
		if page == nil && size == nil {
			db.Model(&outlet).Where("merchant_id = ? ", id).Find(&outlet)
		}
	}


	result := make([]model.Outlet, 0)

	for rows.Next() {
		o := new(model.Outlet)
		err = rows.Scan(
			&o.Id,
			&o.Created,
			&o.CreatedBy,
			&o.Modified,
			&o.ModifiedBy,
			&o.Active,
			&o.IsDeleted,
			&o.Deleted,
			&o.Deleted_by,
			&o.OutletName,
			&o.OutletAddress,
			&o.OutletPhone,
			&o.OutletCity,
			&o.OutletProvince,
			&o.OutletPostalCode,
			&o.OutletLongitude,
			&o.OutletLatitude,
			&o.OutletDay,
			&o.OutletHour,
			&o.MerchantId,
		)

		//merchantEmail := new(model.Merchant)
		//db.Table("merchants").
		//	Select("merchants.merchant_email").
		//	Where("id = ?", o.MerchantId).
		//	First(&merchantEmail)
		//o.MerchantEmail = merchantEmail.MerchantEmail
		//if err != nil {
		//	logrus.Error(err)
		//	return nil
		//}

		result = append(result, *o)
	}

	db.Close()
	return result
}
