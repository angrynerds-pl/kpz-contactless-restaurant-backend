package responses

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	uuid "github.com/satori/go.uuid"
)

type MenuResponse struct {
	MenuId uuid.UUID      `json:"menuId"`
	Foods  []FoodResponse `json:"foods"`
}

func NewMenuRestaurantResponse(m *model.Menu) (*MenuResponse, error) {
	req := &MenuResponse{
		MenuId: m.ID,
		Foods:  []FoodResponse{},
	}

	for _, food := range m.Foods {
		fRsp, err := NewFoodResponse(&food)
		if err != nil {
			return nil, err
		}

		req.Foods = append(req.Foods, *fRsp)
	}

	return req, nil
}

type FoodResponse struct {
	FoodId      string  `json:"food_id"`
	FoodType    string  `json:"food_type"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func NewFoodResponse(f *model.Food) (*FoodResponse, error) {
	return &FoodResponse{
		FoodId:      f.ID.String(),
		FoodType:    f.FoodType,
		Name:        f.Name,
		Description: f.Description,
		Price:       f.Price,
	}, nil
}
