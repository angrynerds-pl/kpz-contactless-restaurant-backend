package model

import uuid "github.com/satori/go.uuid"

type Restaurant struct {
	Base
	CompanyId uuid.UUID
	AddressId uuid.UUID
	Address   Address
	QRCodeId  uuid.UUID
	QrCode    QrCode
	Menus     []Menu

	name        string
	description string
}

type QrCode struct {
	Base
	url  string
	file byte
}
