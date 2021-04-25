package requests

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/labstack/echo/v4"
)

type CreateRestaurantRequest struct {
	Restaurant struct {
		Name        string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`
	} `json:"restaurant"`
}

func (req *CreateRestaurantRequest) Bind(c echo.Context, r *model.Restaurant, a *model.Address) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	r.Name = req.Restaurant.Name
	r.Description = req.Restaurant.Description

	return nil
}

type GetRestaurantRequest struct {
}

func (req *GetRestaurantRequest) Bind(c echo.Context, r *model.Restaurant) error {
	return nil
}
