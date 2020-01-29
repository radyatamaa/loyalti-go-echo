package repository

import (
	"database/sql"
	"fmt"

	"github.com/radyatamaa/loyalti-go-echo/src/database"
	"github.com/radyatamaa/loyalti-go-echo/src/domain/model"
)

const (
	MemberTypeCard 	= "Member"
	GoldTier 		= "Gold"
	SilverTier 		= "Silver"
	PlatinumTier 	= "Platinum"
)

func GetCardMerchant(page *int, size *int, id *int, card_type *string) []model.Card {

	db := database.ConnectionDB()
	//db := database.ConnectPostgre()
	var kartu []model.Card
	var rows *sql.Rows
	var err error
	var total int

	if id != nil {

		if page != nil && size != nil {

			rows, err = db.Where("program_id = ?", id).Find(&kartu).Order("title").Count(total).Limit(*size).Offset(*page).Rows()
			if err != nil {
				panic(err)
			}

			if card_type != nil {
				if page != nil && size != nil && id != nil {
					fmt.Println("card type dapat ga")
					rows, err = db.Where("program_id = ? AND card_type = ?", id, card_type).Find(&kartu).Order("created asc").Count(total).Limit(*size).Offset(*page).Rows()
					if err != nil {
						panic(err)
					}
				}
			}
		}  else {
				rows, err = db.Find(&kartu).Rows()
				if err != nil {
					panic(err)
				}
			}
		}


	result := make([]model.Card, 0)

	for rows.Next() {
		o := new(model.Card)
		err = rows.Scan(
			&o.Id,
			&o.Created,
			&o.CreatedBy,
			&o.Modified,
			&o.ModifiedBy,
			&o.Active,
			&o.IsDeleted,
			&o.Deleted,
			&o.DeletedBy,
			&o.Title,
			&o.Description,
			&o.FontColor,
			&o.TemplateColor,
			&o.IconImage,
			&o.TermsAndCondition,
			&o.Benefit,
			&o.ValidUntil,
			&o.CurrentPoint,
			&o.IsValid,
			&o.ProgramId,
			&o.CardType,
			&o.IconImageStamp,
			&o.MerchantId,
		)
		//add tier
		if o.CardType == MemberTypeCard{
			if o.CurrentPoint >= 0 && o.CurrentPoint <= 100{
				o.Tier = SilverTier
			}else if o.CurrentPoint > 100 && o.CurrentPoint <= 230{
				o.Tier = GoldTier
			}else if o.CurrentPoint > 230 && o.CurrentPoint <= 400{
				o.Tier = PlatinumTier
			}
		}else{
			o.Tier = "Ga ada tier"
		}

		result = append(result, *o)
	}

	db.Close()
	return result
}