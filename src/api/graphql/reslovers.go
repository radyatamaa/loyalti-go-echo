package graphQL

import (
	"github.com/graphql-go/graphql"
	"github.com/radyatamaa/loyalti-go-echo/src/domain"
)

type Song domain.Song

func SongResolver(p graphql.ResolveParams) (interface{}, error) {
	// Strip the name from arguments and assert that it's a string

	///if used parameter
	//name, ok := p.Args["name"].(string)
	//if ok {
	//	users := r.db.GetAll()
	//	return users, nil
	//}
	//

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
