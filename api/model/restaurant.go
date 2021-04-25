package model

import uuid "github.com/satori/go.uuid"

type Restaurant struct {
	Base
	//CompanyId uuid.UUID
	OwnerID uuid.UUID `gorm:"type:uuid"`
	//Address Address   `gorm:"foreignKey:RestaurantId"`
	//QRCodeID  uuid.UUID
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
