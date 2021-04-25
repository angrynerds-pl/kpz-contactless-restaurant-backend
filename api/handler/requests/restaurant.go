package requests

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/labstack/echo/v4"
)

type CreateRestaurantRequest struct {
	Restaurant struct {
		Name        string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`

		Address Address
	} `json:"restaurant"`
}

func (req *CreateRestaurantRequest) Bind(c echo.Context, r *model.Restaurant) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	r.Name = req.Restaurant.Name
	r.Description = req.Restaurant.Description

	r.Address.AddressLine1 = req.Restaurant.Address.AddressLine1
	r.Address.AddressLine2 = req.Restaurant.Address.AddressLine2
	r.Address.City = req.Restaurant.Address.City
	r.Address.PostCode = req.Restaurant.Address.PostCode

	return nil
}
