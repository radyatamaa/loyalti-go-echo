package graphQL

import (
	"fmt"
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/repository"
)

type Song model.Song

func MerchantResolver(p graphql.ResolveParams) (interface{}, error) {
	page, ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	email, top := p.Args["email"].(string)

	if ok && sip && top {
		var pages *int = &page
		var sizes *int = &size
		var emails *string = &email
		merchant := repository.GetMerchant(pages, sizes, emails)
		fmt.Println(merchant)
		return merchant, nil
	}

	merchant := repository.GetMerchant(nil, nil, nil)
	fmt.Println(merchant)
	return merchant, nil
}

func MerchantCategoryResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		category := repository.GetCategory(pages, sizes)
		fmt.Println(category)
		return category,nil
	}

	category := repository.GetCategory(nil,nil)

	return category, nil
}

func MerchantCardResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)

	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
	card := repository.GetCard(pages, sizes)
	fmt.Println(card)
	return card, nil
	}
	card := repository.GetCard(nil,nil)
	return card, nil
}

func SocialMediaResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		sosmed := repository.GetSocialMedia(pages, sizes)
		fmt.Println(sosmed)
		return sosmed,nil
	}
	sosmed := repository.GetSocialMedia(nil, nil)
	return sosmed, nil
}

func ProgramResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,deh := p.Args["sort"].(int)
	if ok && sip && deh{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		program := repository.GetProgram(pages, sizes, sorts)
		fmt.Println(program)
		return program,nil
	}

	program := repository.GetProgram(nil,nil, nil)

	return program, nil
}
//func ProgramByMerchantId(p graphql.ResolveParams) (interface{}, error){
//	page := p.Args["page"].(int)
//	size := p.Args["size"].(int)
//	program := repository.GetProgramByMerchantId(page, size)
//	//fmt.Println(program)
//	return program, nil
//}

//func ProgramResolverByDate(p graphql.ResolveParams) (interface{}, error) {
//	page := p.Args["page"].(int)
//	size := p.Args["size"].(int)
//	sort := p.Args["sort"].(int)
//	program1 := repository.(page, size, sort)
//	fmt.Println(program1)
//	return program1, nil
//}
func OutletResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	if ok && sip {
		var pages *int = &page
		var sizes *int = &size
		outlet := repository.GetOutlet(pages, sizes)
		fmt.Println(outlet)
		return outlet,nil
	}

	outlet := repository.GetOutlet(nil,nil)

	return outlet, nil
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
func MerchantResolve() {

}
