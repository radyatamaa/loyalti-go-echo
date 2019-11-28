package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"time"
)

func main() {
	db := database.ConnectionDB()

	var program1 = model.Program{
		Created:            time.Now(),
		CreatedBy:          "Admin",
		Modified:           time.Now(),
		ModifiedBy:         "Admin",
		Active:             true,
		IsDeleted:          false,
		Deleted:            nil,
		Deleted_by:         "",
		ProgramName:        "Diskon Indomaret",
		ProgramImage:       "",
		ProgramStartDate:   time.Now(),
		ProgramEndDate:     time.Time{},
		ProgramDescription: "Belanja perlengkapan rumah tangga diskon sampai 70%",
		Card:               "Member",
		OutletID:           0,
		MerchantId:         5,
		CategoryId:         2,
	}

	var program2 = model.Program{
		Created:            time.Now(),
		CreatedBy:          "Admin",
		Modified:           time.Now(),
		ModifiedBy:         "Admin",
		Active:             true,
		IsDeleted:          false,
		Deleted:            nil,
		Deleted_by:         "",
		ProgramName:        "Diskon Alfamart",
		ProgramImage:       "",
		ProgramStartDate:   time.Now(),
		ProgramEndDate:     time.Time{},
		ProgramDescription: "Belanja perlengkapan rumah tangga di Alfamart diskon sampai 70%",
		Card:               "Member",
		OutletID:           0,
		MerchantId:         5,
		CategoryId:         2,
	}

	var program3 = model.Program{
		Created:            time.Now(),
		CreatedBy:          "Admin",
		Modified:           time.Now(),
		ModifiedBy:         "Admin",
		Active:             true,
		IsDeleted:          false,
		Deleted:            nil,
		Deleted_by:         "",
		ProgramName:        "Diskon Alfamidi",
		ProgramImage:       "",
		ProgramStartDate:   time.Now(),
		ProgramEndDate:     time.Time{},
		ProgramDescription: "Belanja perlengkapan rumah tangga di alfamidi diskon sampai 90%",
		Card:               "Chop",
		OutletID:           0,
		MerchantId:         6,
		CategoryId:         2,
	}

	var program4 = model.Program{
		Created:            time.Now(),
		CreatedBy:          "Admin",
		Modified:           time.Now(),
		ModifiedBy:         "Admin",
		Active:             true,
		IsDeleted:          false,
		Deleted:            nil,
		Deleted_by:         "",
		ProgramName:        "Promo Kosmetik",
		ProgramImage:       "",
		ProgramStartDate:   time.Now(),
		ProgramEndDate:     time.Time{},
		ProgramDescription: "Belanja perlengkapan kosmetik diskon sampai 70%",
		Card:               "Member",
		OutletID:           0,
		MerchantId:         5,
		CategoryId:         3,
	}

	var program5 = model.Program{
		Created:            time.Now(),
		CreatedBy:          "Admin",
		Modified:           time.Now(),
		ModifiedBy:         "Admin",
		Active:             true,
		IsDeleted:          false,
		Deleted:            nil,
		Deleted_by:         "",
		ProgramName:        "YES",
		ProgramImage:       "",
		ProgramStartDate:   time.Now(),
		ProgramEndDate:     time.Time{},
		ProgramDescription: "Diskon Nike Sport",
		Card:               "Member",
		OutletID:           0,
		MerchantId:         7,
		CategoryId:         1,
	}

	db.Create(&program1)
	db.Create(&program2)
	db.Create(&program3)
	db.Create(&program4)
	db.Create(&program5)
}
