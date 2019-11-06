package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/database"

)
type Product struct {
	Code string
	Price uint
}

type Merchant struct {
	gorm.Model
	MerchantName string
	PhoneNumber string
	Email       string
}
func main() {
	db := database.ConnectionDB()
	// Migrate the schema
	db.AutoMigrate(&Merchant{})
	//if err == nil{
	//	panic("success migration")
	//}

	var merchant []Merchant
	db.Find(&merchant)
	fmt.Println(merchant)
	//// Create
	//db.Create(&Product{Code: "L1212", Price: 1000})
	//
	//// Read
	//var product Product
	//db.First(&product, 1) // find product with id 1
	//db.First(&product, "code = ?", "L1212") // find product with code l1212
	//
	//// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// Delete - delete product
	//db.Delete(&product)
}
