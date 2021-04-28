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

func (h *Handler) CreateRestaurant(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())

	}

	req := &requests.CreateRestaurantRequest{}

	r, err := req.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if err := h.restaurantStore.AddRestaurantToUser(*userId, r); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	rsp, err := responses.NewRestaurantResponse(r)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)

}

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

	rsp, err := responses.NewRestaurantResponse(restaurant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)
}

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

func (h *Handler) UpdateRestaurant(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())
	}

	req := &requests.UpdateRestaurantRequest{}

	r, err := req.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	restaurantId, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return err
	}

	userId, err := userIDFromToken(c)

	r.ID = restaurantId
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	restaurant, err := h.restaurantStore.Update(*userId, r)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	rsp, err := responses.NewRestaurantResponse(restaurant)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)
}

func (h *Handler) RemoveRestaurantFromUser(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())
	}

	var r model.Restaurant
	req := &requests.DeleteRestaurantRequest{}

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
	err = h.restaurantStore.DeleteRestaurantFromUser(*userId, restaurantId)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]string{"userId": userId.String(), "restaurantId": restaurantId.String()})
}

func (h *Handler) AddAddressToRestaurant(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())
	}

	req := &requests.AddAddressToRestaurantRequest{}

	addr, err := req.Bind(c)
	if err != nil {
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
	err = h.restaurantStore.Add(restaurantId, addr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, map[string]string{"userId": userId.String(), "restaurantId": restaurantId.String()})
}
