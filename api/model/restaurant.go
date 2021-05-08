package model

import uuid "github.com/satori/go.uuid"

type Restaurant struct {
	Base
	//CompanyId uuid.UUID
	OwnerID uuid.UUID `gorm:"type:uuid"`
	Address Address   `gorm:"foreignKey:RestaurantID"`
	//QRCodeId  uuid.UUID
	//QrCode    QrCode
	Menu Menu `gorm:"foreignKey:RestaurantID"`

	Name        string
	Description string
}

type QrCode struct {
	Base
	url  string
	file byte
}
