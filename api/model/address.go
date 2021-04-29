package model

import uuid "github.com/satori/go.uuid"

type Address struct {
	Base

	RestaurantID uuid.UUID `gorm:"unique"`

	AddressLine1 string
	AddressLine2 string
	City         string
	PostCode     string
}
