package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao/postgres"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	user *postgres.User
}

func (u User) Create(c echo.Context) error {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	users, err := u.user.Create(&user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (u User) Get(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	us, err := u.user.Get(user.ID.String())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, us)
	//user := c.Get("user").(*jwt.Token)
	//claims := user.Claims.(jwt.MapClaims)
	//name := claims["name"].(string)
	//return c.String(http.StatusOK, "Welcome "+name+"!")

}

func (u User) GetAll(c echo.Context) error {
	panic("implement me")
}

func (u User) Delete(c echo.Context) error {
	panic("implement me")
}

func (u User) Update(c echo.Context) error {
	panic("implement me")
}
