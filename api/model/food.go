package model

import uuid "github.com/satori/go.uuid"

type Food struct {
	Base
	//CuisineTypeId uuid.UUID
	FoodType    string
	Name        string
	Description string
	Price       float64

	OrderID uuid.UUID
}

type Menu struct {
	Base
	Foods []Food `gorm:"many2many:menu_foods;"`

	RestaurantID uuid.UUID
}
