package model

import "time"

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

type Merchant struct {
	GeneralModels
	MerchantName          string `json:"merchant_name"`
	MerchantEmail         string `json:"merchant_email"`
	MerchantPhoneNumber   string `json:"merchant_phone_number"`
	MerchantProvince      string `json:"merchant_province"`
	MerchantCity          string `json:"merchant_city"`
	MerchantAddress       string `json:"merchant_address"`
	MerchantPostalCode    string `json:"merchant_postal_code"`
	MerchantCategoryId    int    `json:"merchant_category_id"`
	MerchantWebsite       string `json:"merchant_website"`
	MerchantMediaSocialId int    `json:"merchant_media_social_id"`
	MerchantDescription   string `json:"merchant_description"`
	MerchantImageProfile  string `json:"merchant_image_profile"`
	MerchantGallery       string `json:"merchant_gallery"`
}

type MerchantCategory struct {
	GeneralModels
	CategoryName string `json:"category_name"`
	ImageUrl     string `json:"image_url"`
}

type MerchantSocialMedia struct {
	GeneralModels
	SocialMediaName     string `json:"social_media_name"`
	SocialMediaImageUrl string `json:"social_media_image_url"`
}

type CardType struct {
	GeneralModels
	CardTypeName string `json:"card_type_name"`
}

type Outlet struct {
	GeneralModels
	OutletName       string    `json:"outlet_name"`
	OutletAddress    string    `json:"outlet_address"`
	OutletPhone      string    `json:"outlet_phone"`
	OutletCity       string    `json:"outlet_city"`
	OutletProvince   string    `json:"outlet_province"`
	OutletPostalCode string    `json:"outlet_postal_code"`
	OutletLongitude  string    `json:"outlet_longitude"`
	OutletLatitude   string    `json:"outlet_latitude"`
	OutletDay        time.Time `json:"outlet_day"`
	OutletHour       time.Time `json:"outlet_hour"`
}

type Program struct {
	GeneralModels
	ProgramName        string    `json:"program_name"`
	ProgramImage       string    `json:"program_image"`
	ProgramStartDate   time.Time `json:"program_start_date"`
	ProgramEndDate     time.Time `json:"program_end_date"`
	ProgramDescription string    `json:"program_description"`
	Card               string    `json:"card"`
}
