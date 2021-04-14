package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao/user"
)

type Handler struct {
	userStore user.Store
}

func NewHandler(us user.Store) *Handler {
	return &Handler{
		userStore: us,
	}
}
