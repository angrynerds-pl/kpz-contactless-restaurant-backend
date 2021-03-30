package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao/postgres"
	"github.com/labstack/echo/v4"
)

const (
	//Models
	user         = "user"

)

type CRUDHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetAll(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}


type Handler struct {
	User dao.User
}

//NewHandler returns a new BaseHandler
func NewCRUDHandler(name string) CRUDHandler {
	switch name {
	case user:
		return &User{user: &postgres.User{Db: }}
	}
	return nil
}
