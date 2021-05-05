package model

type Food struct {
	Base
	//CuisineTypeId uuid.UUID
	FoodType    string
	Name        string `gorm:"unique"`
	Description string
	Price       string
}

type Menu struct {
	Base
	Foods []Food `gorm:"many2many:menu_foods;"`
}
