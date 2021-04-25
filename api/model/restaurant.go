package model

import uuid "github.com/satori/go.uuid"

type Restaurant struct {
	Base
	//CompanyId uuid.UUID
	OwnerId uuid.UUID `gorm:"type:uuid"`
	Address Address   `gorm:"foreignKey:RestaurantId"`
	//QRCodeId  uuid.UUID
	//QrCode    QrCode
	//Menus     []Menu

	Name        string
	Description string
}

type QrCode struct {
	Base
	url  string
	file byte
}
