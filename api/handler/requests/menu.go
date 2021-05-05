package requests

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/labstack/echo/v4"
)

type AddFoodToMenuRequest struct {
	Food struct {
		Name        string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`
		FoodType    string `param:"foodType" query:"foodType" form:"foodType" json:"foodType" xml:"foodType" validate:"required"`
		Price       string `param:"price" query:"price" form:"price" json:"price" xml:"price" validate:"required"`
	} `json:"food"`
}

func (req AddFoodToMenuRequest) Bind(c echo.Context) (*model.Food, error) {
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	if err := c.Validate(req); err != nil {
		return nil, err
	}

	return &model.Food{
		FoodType:    req.Food.FoodType,
		Name:        req.Food.Name,
		Description: req.Food.Description,
		Price:       req.Food.Price,
	}, nil
}

type RemoveFoodToMenuRequest struct {
	Food struct {
		Name string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
	} `json:"food"`
}

func (req RemoveFoodToMenuRequest) Bind(c echo.Context) (*model.Food, error) {
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	if err := c.Validate(req); err != nil {
		return nil, err
	}

	return &model.Food{
		Name: req.Food.Name,
	}, nil
}
