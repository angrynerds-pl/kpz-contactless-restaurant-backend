package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Order struct {
	Base

	Foods  []Food
	Status status

	OrderRestaurantCustomer []OrderRestaurantCustomer
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.Status = New
	return
}

type OrderRestaurantCustomer struct {
	OrderID      uuid.UUID `gorm:"type:uuid"`
	RestaurantID uuid.UUID `gorm:"type:uuid"`
	CustomerID   uuid.UUID `gorm:"type:uuid"`
}

type status int

const (
	New status = iota
	Preparing
	Prepared
)
