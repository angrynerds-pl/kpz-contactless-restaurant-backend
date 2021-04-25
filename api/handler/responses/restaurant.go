package responses

import "github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"

type restaurantResponse struct {
	Restaurant struct {
		Name        string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`

		Address Address
	} `json:"restaurant"`
}

func NewRestaurantResponse(r *model.Restaurant) (*restaurantResponse, error) {
	req := new(restaurantResponse)
	req.Restaurant.Name = r.Name
	req.Restaurant.Description = r.Description

	req.Restaurant.Address.AddressLine1 = r.Address.AddressLine1
	req.Restaurant.Address.AddressLine2 = r.Address.AddressLine2
	req.Restaurant.Address.City = r.Address.City
	req.Restaurant.Address.PostCode = r.Address.PostCode

	return req, nil
}
