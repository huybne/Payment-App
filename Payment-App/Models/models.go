package Models

import (
	"gorm.io/gorm"
)

type PaymentDetail struct {
	gorm.Model
	CardOwnerName  string `gorm:"type:varchar(50)" json:"card_owner_name"`
	CardNumber     string `gorm:"type:varchar(16)" json:"cardNumber"`
	ExpirationDate string `gorm:"type:varchar(5)" json:"expirationDate"`
	SecurityCode   string `gorm:"type:varchar(3)" json:"securityCode"`
}

type Payments struct {
	Payments []*PaymentDetail
}
