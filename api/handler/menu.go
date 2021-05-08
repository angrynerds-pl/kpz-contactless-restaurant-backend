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

// AddFoodToMenu
// @Summary AddFoodToMenu Adds Food to Menu using RestaurantId and Food
// @Description AddFoodToMenu
// @Tags Restaurants, Menus
// @Accept  json
// @Produce  json
// @Param id path string true "Id of restaurant"
// @Param restaurant body requests.AddFoodToMenuRequest true "Details of food"
// @Success 200 {object} responses.MenuResponse
// @Failure default {object} utils.Error
// @Router /users/restaurants/:id/menus/ [post]
func (h *Handler) AddFoodToMenu(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())
	}

	req := requests.AddFoodToMenuRequest{}

	f, err := req.Bind(c)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	restaurantId, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	restaurant, err := h.restaurantStore.GetByID(*userId, restaurantId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if err := h.menuStore.AddFoodToMenu(&restaurant.Menu, f); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	rsp, err := responses.NewMenuRestaurantResponse(&restaurant.Menu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)

}

// GetMenu
// @Summary GetMenu Get menu using restaurantId
// @Description GetMenu
// @Tags Restaurants, Menus
// @Accept  json
// @Produce  json
// @Param id path string true "Id of restaurant"
// @Success 200 {object} responses.MenuResponse
// @Failure default {object} utils.Error
// @Router /users/restaurants/:id/menus/ [get]
func (h Handler) GetMenu(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())
	}

	restaurantId, err := uuid.FromString(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	restaurant, err := h.restaurantStore.GetByID(*userId, restaurantId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	menu, err := h.menuStore.GetMenu(restaurant.Menu.ID)

	rsp, err := responses.NewMenuRestaurantResponse(menu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)
}

// RemoveFoodFromMenu
// @Summary RemoveFoodFromMenu Removes food from Menu using restaurantId and foodId
// @Description RemoveFoodFromMenu
// @Tags Restaurants, Menus
// @Accept  json
// @Produce  json
// @Param id path string true "Id of restaurant"
// @Param foodId path string true "Id of food"
// @Success 200 {object} responses.MenuResponse
// @Failure default {object} utils.Error
// @Router /users/restaurants/:id/menus/:foodId/ [delete]
func (h Handler) RemoveFoodFromMenu(c echo.Context) error {
	role, err := userRoleFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if role != model.Owner {
		return c.JSON(http.StatusInternalServerError, utils.AccessForbidden())
	}

	restaurantId, err := uuid.FromString(c.Param("id"))
	foodId, err := uuid.FromString(c.Param("foodId"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	restaurant, err := h.restaurantStore.GetByID(*userId, restaurantId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	if h.menuStore.RemoveFoodFromMenu(restaurant.Menu.ID, foodId) != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	menu, err := h.menuStore.GetMenu(restaurant.Menu.ID)

	rsp, err := responses.NewMenuRestaurantResponse(menu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)
}
