package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func userIDFromToken(c echo.Context) (*uuid.UUID, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := claims["id"].(string)

	idUuid, err := uuid.FromString(id)
	if err != nil {
		return nil, err
	}

	return &idUuid, nil
}

func userRoleFromToken(c echo.Context) (model.RoleName, error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	role := claims["role"].(float64)

	return model.ToRoleName(role), nil
}
