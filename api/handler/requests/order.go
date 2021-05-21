package requests

import (
	"github.com/labstack/echo/v4"
)

type OrderFoodsRequest struct {
	Order struct {
		FoodIds []string `param:"foodIds" query:"foodIds" form:"foodIds" json:"foodIds" xml:"foodIds" `
	} `json:"order"`
	RestaurantId string `json:"restaurant_id"`
}

func (req *OrderFoodsRequest) Bind(c echo.Context) ([]string, *string, error) {
	if err := c.Bind(req); err != nil {
		return nil, nil, err
	}
	//if err := c.Validate(req); err != nil {
	//	return nil, err
	//}

	var foodIds []string

	for _, id := range req.Order.FoodIds {
		foodIds = append(foodIds, id)
	}

	return foodIds, &req.RestaurantId, nil

}
