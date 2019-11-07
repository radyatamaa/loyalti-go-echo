package domain

type Song struct {
	ID       string `json:"id,omitempty"`
	Album    string `json:"album"`
	Title    string `json:"title"`
	Duration string `json:"duration"`
	Type     string `json:"type"`
}

type Merchant struct {
	GeneralModels
	MerchantName string `json:"merchant_name"`
	MerchantEmail string  `json:"merchant_email"`
	MerchantPhoneNumber string  `json:"merchant_phone_number"`
	MerchantProvince string `json:"merchant_province"`
	MerchantCity string `json:"merchant_city"`
	MerchantAddress string `json:"merchant_address"`
	MerchantPostalCode string `json:"merchant_postal_code"`
	MerchantCategory string `json:"merchant_category"`
	MerchantWebsite string `json:"merchant_website"`
	MerchantMediaSocial string `json:"merchant_media_social"`
	MerchantDescription string `json:"merchant_description"`
	MerchantImageProfile string `json:"merchant_image_profile"`
	MerchantGallery string `json:"merchant_gallery"`
}

type Program struct {
}

