package responses

import "github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"

type RestaurantResponse struct {
	Restaurant struct {
		Name        string `param:"name" query:"name" form:"name" json:"name" xml:"name" validate:"required"`
		Description string `param:"description" query:"description" form:"description" json:"description" xml:"description" validate:""`
	} `json:"restaurant"`
}

func NewRestaurantResponse(r *model.Restaurant) (*RestaurantResponse, error) {
	req := new(RestaurantResponse)
	req.Restaurant.Name = r.Name
	req.Restaurant.Description = r.Description

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
