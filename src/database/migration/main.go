package main

import (
	"fmt"
	"github.com/beevik/guid"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"time"
)

func main() {
	//var id = 16
	//var pay = 7000
	//repository.TotalPoint(&id, &pay)

	//guids := guid.New()
	//fmt.Println(guids)


	db := database.ConnectionDB()
	//db.AutoMigrate(model.Card{})
	fmt.Println("migrasi berhasil")

	 card := model.Card{
		Id:                guid.Guid{},
		Created:           time.Time{},
		CreatedBy:         "Admin",
		Modified:          time.Time{},
		ModifiedBy:        "Admin",
		Active:            true,
		IsDeleted:         false,
		Deleted:           nil,
		DeletedBy:         "",
		Title:             "lalala",
		Description:       "deskripsi",
		FontColor:         "hitam",
		TemplateColor:     "batik",
		IconImage:         "www.goggle.com",
		TermsAndCondition: "TnC",
		Benefit:           "untung",
		ValidUntil:        time.Time{},
		RewardTarget:      "",
		IsValid:           false,
		ProgramId:         0,
		CardType:          "3",
		IconImageStamp:    "",
		MerchantId:        0,
	}
	db.Create(&card)
}

