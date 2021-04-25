package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler/requests"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler/responses"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) CreateRestaurant(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())

	}

	var r model.Restaurant
	req := &requests.CreateRestaurantRequest{}

	if err := req.Bind(c, &r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if err := h.restaurantStore.AddRestaurantToUser(*userId, &r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	rsp, err := responses.NewRestaurantResponse(&r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)

}

func (h *Handler) GetRestaurant(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) Restaurants(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) UpdateRestaurant(c echo.Context) error {
	panic("implement me")
}

func (h *Handler) RemoveRestaurantFromUser(c echo.Context) error {
	panic("implement me")
}
