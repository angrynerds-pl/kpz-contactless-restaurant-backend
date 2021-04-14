package model

import uuid "github.com/satori/go.uuid"

type Food struct {
	Base
	CuisineTypeId uuid.UUID
	FoodType      string
	Name          string
	Description   string
	Price         string
}

type Menu struct {
	Base
	Food []Food `gorm:"many2many:menu_foods;"`
}
