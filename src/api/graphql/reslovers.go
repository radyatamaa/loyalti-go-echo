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
	sort, tap := p.Args["sort"].(int)
	if ok && sip && top && tap{
		var pages *int = &page
		var sizes *int = &size
		var emails *string = &email
		var sorts *int = &sort
		merchant := repository.GetMerchant(pages, sizes, sorts, emails)
		fmt.Println(merchant)
		return merchant, nil
	}

	merchant := repository.GetMerchant(nil, nil, nil, nil)
	fmt.Println(merchant)
	return merchant, nil
}

func MerchantCategoryResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,top := p.Args["sort"].(int)
	if ok && sip && top{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		category := repository.GetCategory(pages, sizes, sorts)
		fmt.Println(category)
		return category,nil
	}

	category := repository.GetCategory(nil,nil, nil)

	return category, nil
}

func MerchantCardResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size, sip := p.Args["size"].(int)
	sort, top := p.Args["sort"].(int)

	if ok && sip && top{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
	card := repository.GetCard(pages, sizes, sorts)
	fmt.Println(card)
	return card, nil
	}
	card := repository.GetCard(nil,nil, nil)
	return card, nil
}

func SocialMediaResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,top := p.Args["sort"].(int)
	if ok && sip && top{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		sosmed := repository.GetSocialMedia(pages, sizes, sorts)
		fmt.Println(sosmed)
		return sosmed,nil
	}
	sosmed := repository.GetSocialMedia(nil, nil, nil)
	return sosmed, nil
}

func SpecialProgramResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,deh := p.Args["sort"].(int)
	if ok && sip && deh{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		special := repository.GetSpecialProgram(pages, sizes, sorts)
		fmt.Println(special)
		return special,nil
	}

	special := repository.GetSpecialProgram(nil,nil, nil)

	return special, nil
}

//program function
func ProgramResolver(p graphql.ResolveParams) (interface{}, error) {
	page,ok := p.Args["page"].(int)
	size,sip := p.Args["size"].(int)
	sort,deh := p.Args["sort"].(int)
	category, cat := p.Args["category"].(int)
	if ok && sip && deh && cat {
		fmt.Println("12345")
		var pagination *int = &page
		var sizing *int = &size
		var category *int = &category
		var sorting *int = &sort
		program := repository.GetProgram(pagination,sizing,sorting,category)
		//fmt.Println(program)
		return program,nil
	}else if ok && sip && deh {
		fmt.Println("5678")
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		//var cats *int = &category
		program := repository.GetProgram(pages, sizes, sorts,nil)
		//fmt.Println(program)
		return program,nil
	} else if ok && sip {
		fmt.Println("7890")
		var paging *int = &page
		var sizing *int = &size
		program := repository.GetProgram(paging, sizing, nil, nil)
		//fmt.Println(program)
		return program,nil
	}

	program := repository.GetProgram(nil,nil, nil, nil)

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
	sort, top := p.Args["sort"].(int)
	if ok && sip && top{
		var pages *int = &page
		var sizes *int = &size
		var sorts *int = &sort
		outlet := repository.GetOutlet(pages, sizes, sorts)
		fmt.Println(outlet)
		return outlet,nil
	}

	outlet := repository.GetOutlet(nil,nil, nil)

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
