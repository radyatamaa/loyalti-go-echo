package main

import (
	"github.com/beevik/guid"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"

	//"github.com/radyatamaa/loyalti-go-echo/src/database"

)

func main() {
	db := database.ConnectionDB()
	employee := model.Card{
		Id:                guid.NewString(),
		Created:           time.Now(),
		CreatedBy:         "Admin",
		Modified:          time.Now(),
		ModifiedBy:        "Admin",
		Active:            true,
		IsDeleted:         false,
		Deleted:           nil,
		DeletedBy:         "",
		Title:             "February Blast",
		Description:       "Redeem you point to get FREE Lipstick & Mascara",
		FontColor:         "Black",
		TemplateColor:     "https://loyalti-storage.s3-ap-southeast-1.amazonaws.com/card/pointcard.png",
		IconImage:         "https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTroxoLpVqjJbxzunhPdgB0l4geUC0f_TvDGRRFz_x4gR1vST5Y",
		TermsAndCondition: "TERM AND CONDITION",
		Benefit:           "FREE Lipstick & Mascara",
		ValidUntil:        time.Time{},
		CurrentPoint:      100,
		IsValid:           true,
		ProgramId:         5,
		CardType:          "Voucher",
		IconImageStamp:    "Stamp",
		MerchantId:        5,
	}
	db.Create(&employee)
	//db.AutoMigrate(&employee)
	}



