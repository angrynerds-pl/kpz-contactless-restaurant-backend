package requests

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/labstack/echo/v4"
)

type CreateRestaurantRequest struct {
	Restaurant struct {
		Name        string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`

		Address Address `param:"address" query:"address" form:"address" json:"address" xml:"address" validate:""`
	} `json:"restaurant"`
}

func (req *CreateRestaurantRequest) Bind(c echo.Context) (*model.Restaurant, error) {
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	if err := c.Validate(req); err != nil {
		return nil, err
	}
	r := model.Restaurant{}

	r.Name = req.Restaurant.Name
	r.Description = req.Restaurant.Description

	r.Address.AddressLine1 = req.Restaurant.Address.AddressLine1
	r.Address.AddressLine2 = req.Restaurant.Address.AddressLine2
	r.Address.City = req.Restaurant.Address.City
	r.Address.PostCode = req.Restaurant.Address.PostCode

	return &r, nil
}

type GetRestaurantRequest struct {
}

func (req *GetRestaurantRequest) Bind(c echo.Context, r *model.Restaurant) error {
	return nil
}

type UpdateRestaurantRequest struct {
	Restaurant struct {
		Name        string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`

		Address Address `param:"address" query:"address" form:"address" json:"address" xml:"address" validate:""`
	} `json:"restaurant"`
}

func (req *UpdateRestaurantRequest) Bind(c echo.Context) (*model.Restaurant, error) {
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	if err := c.Validate(req); err != nil {
		return nil, err
	}
	r := model.Restaurant{}

	r.Name = req.Restaurant.Name
	r.Description = req.Restaurant.Description

	r.Address.AddressLine1 = req.Restaurant.Address.AddressLine1
	r.Address.AddressLine2 = req.Restaurant.Address.AddressLine2
	r.Address.City = req.Restaurant.Address.City
	r.Address.PostCode = req.Restaurant.Address.PostCode

	return &r, nil
}

type AddAddressToRestaurantRequest struct {
	Address Address `param:"address" query:"address" form:"address" json:"address" xml:"address" validate:""`
}

func (req *AddAddressToRestaurantRequest) Bind(c echo.Context) (*model.Address, error) {
	if err := c.Bind(req); err != nil {
		return nil, err
	}
	if err := c.Validate(req); err != nil {
		return nil, err
	}
	a := model.Address{}

	a.AddressLine1 = req.Address.AddressLine1
	a.AddressLine2 = req.Address.AddressLine2
	a.City = req.Address.City
	a.PostCode = req.Address.PostCode

	return &a, nil
}

type DeleteRestaurantRequest struct {
}

func (req *DeleteRestaurantRequest) Bind(c echo.Context, r *model.Restaurant) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	return nil
}
