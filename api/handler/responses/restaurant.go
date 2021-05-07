package responses

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
)

type RestaurantResponse struct {
	Restaurant struct {
		Name        string       `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string       `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`
		Menu        MenuResponse `json:"menu"`
	} `json:"restaurant"`
}

func NewRestaurantResponse(r *model.Restaurant) (*RestaurantResponse, error) {
	req := new(RestaurantResponse)
	req.Restaurant.Name = r.Name
	req.Restaurant.Description = r.Description

	menu, err := NewMenuRestaurantResponse(&r.Menu)
	if err != nil {
		return nil, err
	}

	req.Restaurant.Menu = *menu

	return req, nil
}

type RestaurantsResponse struct {
	Restaurants []RestaurantResponse `json:"restaurants"`
}

func NewRestaurantsResponse(rs []model.Restaurant) (*RestaurantsResponse, error) {
	rsp := &RestaurantsResponse{Restaurants: []RestaurantResponse{}}

	for _, r := range rs {
		restaurantRsp, err := NewRestaurantResponse(&r)
		if err != nil {
			return nil, err
		}
		rsp.Restaurants = append(rsp.Restaurants, *restaurantRsp)
	}
	return rsp, nil
}

type UserRestaurantResponse struct {
	RestaurantID string
	UserID       string
}

func NewUserRestaurantResponse(userId, restaurantId string) (*UserRestaurantResponse, error) {
	return &UserRestaurantResponse{
		RestaurantID: restaurantId,
		UserID:       userId,
	}, nil
}

type RestaurantAddressResponse struct {
	RestaurantID string
	AddressID    string
}

func NewRestaurantAddressResponse(restaurantId, addressId string) (*RestaurantAddressResponse, error) {
	return &RestaurantAddressResponse{
		RestaurantID: restaurantId,
		AddressID:    addressId,
	}, nil
}
