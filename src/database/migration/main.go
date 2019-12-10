package main

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/repository"
)

func main() {
	var id = 16
	var pay = 7000
	repository.TotalPoint(&id, &pay)

}
