package graphQL

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/radyatamaa/loyalti-go-echo/src/domain"
	"github.com/radyatamaa/loyalti-go-echo/src/infrastructure/presistance"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Song domain.Song


func MerchantResolver(p graphql.ResolveParams)(interface{},error)  {
	merchant := presistance.GetAll()
	fmt.Println(merchant)
	return merchant,nil

}
func SongResolver(p graphql.ResolveParams) (interface{}, error) {
	users := []Song{
		Song{
			ID:       "1",
			Album:    "ts-fearless",
			Title:    "Fearless",
			Duration: "4:01",
			Type:     "song",
		},
		Song{
			ID:       "2",
			Album:    "ts-fearless",
			Title:    "Fifteen",
			Duration: "4:54",
			Type:     "song",
		},
	}
	return users, nil
	//return nil, nil
}
func MerchantResolve()  {
	
}

