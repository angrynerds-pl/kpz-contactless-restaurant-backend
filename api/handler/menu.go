package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler/requests"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler/responses"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// AddFoodToMenu
// @Summary AddFoodToMenu
// @Description AddFoodToMenu
// @Tags Restaurants, Menus
// @Accept  json
// @Produce  json
// @Param restaurant body requests.AddFoodToMenuRequest true "Details of food"
// @Success 200 {object} responses.AddFoodToMenuResponse
// @Failure default {object} utils.Error
// @Router /menus/ [post]
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

	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if err := h.menuStore.AddFoodToMenu(*userId, f); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	rsp := responses.FoodToMenuResponse{
		MenuId: menuId,
		Food: struct {
			Name string
		}{Name: req.Food.Name},
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, rsp)

}

func (h Handler) GetMenu(c echo.Context) error {
	return nil
}

func (h Handler) RemoveFoodFromMenu(c echo.Context) error {
	return nil
}
