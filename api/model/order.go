package model

import uuid "github.com/satori/go.uuid"

type Order struct {
	Base

	Foods  []Food
	Status string

	OrderRestaurantCustomer []OrderRestaurantCustomer
}

type OrderRestaurantCustomer struct {
	OrderID      uuid.UUID `gorm:"type:uuid"`
	RestaurantID uuid.UUID `gorm:"type:uuid"`
	CustomerID   uuid.UUID `gorm:"type:uuid"`
}
