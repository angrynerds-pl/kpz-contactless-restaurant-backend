package model

import uuid "github.com/satori/go.uuid"

type Address struct {
	Base

	RestaurantId uuid.UUID

	AddressLine1 string
	AddressLine2 string
	City         string
	PostCode     string
}
