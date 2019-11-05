package domain

import (
	"github.com/jinzhu/gorm"
)

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

type Merchant struct {
	gorm.Model
	MerchantName string
	PhoneNumber string
	Email       string
}

