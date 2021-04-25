package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler/requests"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler/responses"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/utils"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// CreateRestaurant
// @Summary Create Restaurant
// @Description Create Restaurant
// @Tags restaurants
// @Accept  json
// @Produce  json
// @Param restaurant body requests.CreateRestaurantRequest true "Details of restaurant"
// @Success 200 {object} responses.RestaurantResponse
// @Failure default {object} utils.Error
// @Router /users/restaurants/ [post]
func (h *Handler) CreateRestaurant(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())

	}

	var r model.Restaurant
	var a model.Address
	req := &requests.CreateRestaurantRequest{}

	if err := req.Bind(c, &r, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if err := h.restaurantStore.AddRestaurantToUser(*userId, r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	rsp, err := responses.NewRestaurantResponse(&r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)

}

// GetRestaurant
// @Summary Get Restaurant by id
// @Description Get Restaurant by id
// @Tags restaurants
// @Accept  json
// @Produce  json
// @Param id path string true "Id of restaurant"
// @Success 200 {object} responses.RestaurantResponse
// @Failure default {object} utils.Error
// @Router /users/restaurants/{id} [get]
func (h *Handler) GetRestaurant(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())
	}

	var r model.Restaurant
	req := &requests.GetRestaurantRequest{}

	if err := req.Bind(c, &r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	restaurantId, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return err
	}

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	restaurant, err := h.restaurantStore.GetByID(*userId, restaurantId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if restaurant == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}

	rsp, err := responses.NewRestaurantResponse(restaurant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)
}

// Restaurant
// @Summary Get all Restaurant
// @Description Get all Restaurant
// @Tags restaurants
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.RestaurantResponse
// @Failure default {object} utils.Error
// @Router /users/restaurants/ [get]
func (h *Handler) Restaurants(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())
	}

	var r model.Restaurant
	req := &requests.GetRestaurantRequest{}

	if err := req.Bind(c, &r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	restaurants, err := h.restaurantStore.GetAll(*userId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	rsp, err := responses.NewRestaurantsResponse(restaurants)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)
}

// UpdateRestaurant
// @Summary Update Restaurant
// @Description Update Restaurant
// @Tags restaurants
// @Accept  json
// @Produce  json
// @Param id path string true "Id of restaurant"
// @Param restaurant body requests.UpdateRestaurantRequest true "Details of restaurant"
// @Success 200 {object} responses.RestaurantResponse
// @Failure default {object} utils.Error
// @Router /users/restaurants/{id} [put]
func (h *Handler) UpdateRestaurant(c echo.Context) error {
	panic("implement me")
}

// RemoveRestaurantFromUser
// @Summary Remove restaurant from user
// @Description Remove restaurant from user
// @Tags restaurants
// @Accept  json
// @Produce  json
// @Param id path string true "Id of restaurant"
// @Success 200 {object} responses.RestaurantResponse
// @Failure default {object} utils.Error
// @Router /users/restaurants/{id} [delete]
func (h *Handler) RemoveRestaurantFromUser(c echo.Context) error {
	panic("implement me")
}
