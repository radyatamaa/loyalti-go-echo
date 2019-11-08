package graphQL

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/model"
	"github.com/felixsiburian/loyalti-go-echo/src/domain/repository"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Song model.Song


func MerchantResolver(p graphql.ResolveParams)(interface{},error)  {
	merchant := repository.GetAll()
	fmt.Println(merchant)
	return merchant,nil
}

func MerchantCardResolver(p graphql.ResolveParams)(interface{}, error){
		card := repository.GetAll()
		fmt.Println(card)
		return card, nil
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

