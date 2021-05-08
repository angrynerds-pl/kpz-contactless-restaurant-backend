package requests

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/labstack/echo/v4"
)

type AddFoodToMenuRequest struct {
	Food struct {
		Name        string  `param:"name" query:"name" form:"name" json:"name" xml:"name" `
		Description string  `param:"description" query:"description" form:"description" json:"description" xml:"description"`
		FoodType    string  `param:"food_type" query:"food_type" form:"food_type" json:"food_type" xml:"food_type"`
		Price       float64 `param:"price" query:"price" form:"price" json:"price" xml:"price" `
	} `json:"food"`
}

func (req *AddFoodToMenuRequest) Bind(c echo.Context) (*model.Food, error) {
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	//if err := c.Validate(req); err != nil {
	//	return nil, err
	//}

	return &model.Food{
		Name:        req.Food.Name,
		Description: req.Food.Description,
		FoodType:    req.Food.FoodType,
		Price:       req.Food.Price,
	}, nil
}

//type RemoveFoodToMenuRequest struct {
//	Food struct {
//		Id string `param:"id" query:"id" form:"id" json:"id" xml:"id" validate:"required"`
//	} `json:"food"`
//}
//
//func (req *RemoveFoodToMenuRequest) Bind(c echo.Context) (*model.Food, error) {
//	if err := c.Bind(req); err != nil {
//		return nil, err
//	}
//	if err := c.Validate(req); err != nil {
//		return nil, err
//	}
//
//	id,err := uuid.FromString(req.Food.Id)
//	if err != nil {
//		return nil, err
//	}
//
//	return &model.Food{Base: model.Base{ID: id}}, nil
//}
