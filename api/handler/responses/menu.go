package responses

import uuid "github.com/satori/go.uuid"

type FoodToMenuResponse struct {
	MenuId uuid.UUID
	Food   struct {
		Name string
	}
}

type FoodResponse struct {
}
