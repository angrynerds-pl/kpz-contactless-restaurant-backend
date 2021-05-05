package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao/menu"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao/restaurant"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao/user"
)

type Handler struct {
	userStore       user.Store
	restaurantStore restaurant.Store
	menuStore       menu.Menu
}

func NewHandler(us user.Store, rs restaurant.Store, ms menu.Menu) *Handler {
	return &Handler{
		userStore:       us,
		restaurantStore: rs,
		menuStore:       ms,
	}
}
