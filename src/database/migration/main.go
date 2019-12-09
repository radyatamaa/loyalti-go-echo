package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func main() {

	db := database.ConnectionDB()

	program := model.Program{
		Created:            time.Now(),
		CreatedBy:          "Admin",
		Modified:           time.Now(),
		ModifiedBy:         "Admin",
		Active:             true,
		IsDeleted:          false,
		Deleted:            nil,
		Deleted_by:         "",
		ProgramName:        "Diskon 70% Goride",
		ProgramImage:       "",
		ProgramStartDate:   time.Now(),
		ProgramEndDate:     time.Now(),
		ProgramDescription: "Diskon Goride sampai dengan 70% maksimal 10k",
		Card:               "Member",
		OutletID:           "1",
		MerchantId:         9,
		CategoryId:         2,
		Benefit:            "Tidak ada",
		TermsAndCondition:  "Tidak ada",
		Tier:               "Platinum",
		RedeemRules:        "Tidak ada",
		RewardTarget:       0,
		QRCodeId:           "1",
	}
	db.Create(&program)
}
