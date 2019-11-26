package model

import (
	"time"
)

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
	Id           int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created      time.Time  `json:"created"`
	CreatedBy    string     `json:"created_by"`
	Modified     time.Time  `json:"modified"`
	ModifiedBy   string     `json:"modified_by"`
	Active       bool       `json:"active"`
	IsDeleted    bool       `json:"is_deleted"`
	Deleted      *time.Time `json:"deleted"`
	Deleted_by   string     `json:"deleted_by"`
	CategoryName string     `json:"category_name"`
	ImageUrl     string     `json:"image_url"`
}

type MerchantSocialMedia struct {
	Id                  int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created             time.Time  `json:"created"`
	CreatedBy           string     `json:"created_by"`
	Modified            time.Time  `json:"modified"`
	ModifiedBy          string     `json:"modified_by"`
	Active              bool       `json:"active"`
	IsDeleted           bool       `json:"is_deleted"`
	Deleted             *time.Time `json:"deleted"`
	Deleted_by          string     `json:"deleted_by"`
	SocialMediaName     string     `json:"social_media_name"`
	SocialMediaImageUrl string     `json:"social_media_image_url"`
}

type CardType struct {
	Id           int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created      time.Time  `json:"created"`
	CreatedBy    string     `json:"created_by"`
	Modified     time.Time  `json:"modified"`
	ModifiedBy   string     `json:"modified_by"`
	Active       bool       `json:"active"`
	IsDeleted    bool       `json:"is_deleted"`
	Deleted      *time.Time `json:"deleted"`
	Deleted_by   string     `json:"deleted_by"`
	CardTypeName string     `json:"card_type_name"`
}

type Outlet struct {
	Id               int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created          time.Time  `json:"created"`
	CreatedBy        string     `json:"created_by"`
	Modified         time.Time  `json:"modified"`
	ModifiedBy       string     `json:"modified_by"`
	Active           bool       `json:"active"`
	IsDeleted        bool       `json:"is_deleted"`
	Deleted          *time.Time `json:"deleted"`
	Deleted_by       string     `json:"deleted_by"`
	OutletName       string     `json:"outlet_name"`
	OutletAddress    string     `json:"outlet_address"`
	OutletPhone      string     `json:"outlet_phone"`
	OutletCity       string     `json:"outlet_city"`
	OutletProvince   string     `json:"outlet_province"`
	OutletPostalCode string     `json:"outlet_postal_code"`
	OutletLongitude  string     `json:"outlet_longitude"`
	OutletLatitude   string     `json:"outlet_latitude"`
	OutletDay        time.Time  `json:"outlet_day"`
	OutletHour       time.Time  `json:"outlet_hour"`
	MerchantId       int        `json:"merchant_id"`
}

type Program struct {
	//GeneralModels
	Id                 int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"`
	Created            time.Time  `json:"created"`
	CreatedBy          string     `json:"created_by"`
	Modified           time.Time  `json:"modified"`
	ModifiedBy         string     `json:"modified_by"`
	Active             bool       `json:"active"`
	IsDeleted          bool       `json:"is_deleted"`
	Deleted            *time.Time `json:"deleted"`
	Deleted_by         string     `json:"deleted_by"`
	ProgramName        string     `json:"program_name"`
	ProgramImage       string     `json:"program_image"`
	ProgramStartDate   time.Time  `json:"program_start_date"`
	ProgramEndDate     time.Time  `json:"program_end_date"`
	ProgramDescription string     `json:"program_description"`
	Card               string     `json:"card"`
	OutletID           int        `json:"outlet_id"`
	MerchantId         int        `json:"merchant_id"`
	MerchantName       string     `json:"merchant_name"`
	CategoryId         int        `json:"category_id"`
}

type SpecialProgram struct {
	Id                 int        `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT NULL"; json:"id"'`
	Created            time.Time  `json:"created"`
	CreatedBy          string     `json:"created_by"`
	Modified           time.Time  `json:"modified"`
	ModifiedBy         string     `json:"modified_by"`
	Active             bool       `json:"active"`
	IsDeleted          bool       `json:"is_deleted"`
	Deleted            *time.Time `json:"deleted"`
	Deleted_by         string     `json:"deleted_by"`
	ProgramName        string     `json:"program_name"`
	ProgramImage       string     `json:"program_image"`
	ProgramStartDate   time.Time  `json:"program_start_date"`
	ProgramEndDate     time.Time  `json:"program_end_date"`
	ProgramDescription string     `json:"program_description"`
	Card               string     `json:"card"`
	OutletID           int        `json:"outlet_id"`
	MerchantId         int        `json:"merchant_id"`
	//MerchantName 		string      `json:"merchant_name"`
}