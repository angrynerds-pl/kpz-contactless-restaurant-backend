package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// Signup
// @Summary Register new user with owner role
// @Description Register new user in server
// @Tags Guest user
// @Accept  json
// @Produce  json
// @Param User body userRegisterRequest true "User credentials"
// @Success 200 {object} userResponse
// @Failure default {object} utils.Error
// @Router /users [post]
// @Decaprated
func (h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	rsp, err := newUserResponse(&u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, rsp)
}

// Signup
// @Summary Register new user with customer role
// @Description Register new user in server
// @Tags Guest user
// @Accept  json
// @Produce  json
// @Param User body userRegisterRequest true "User credentials"
// @Success 200 {object} userResponse
// @Failure default {object} utils.Error
// @Router /auth/owner [post]
func (h *Handler) SignUpCustomer(c echo.Context) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	rsp, err := newUserResponse(&u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, rsp)
}

// Signup
// @Summary Register new user with owner role
// @Description Register new user in server
// @Tags Guest user
// @Accept  json
// @Produce  json
// @Param User body userRegisterRequest true "User credentials"
// @Success 200 {object} userResponse
// @Failure default {object} utils.Error
// @Router /auth/owner [post]
func (h *Handler) SignUpOwner(c echo.Context) error {
	var u model.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	rsp, err := newUserResponse(&u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, rsp)
}

// Login
// @Summary Login to service
// @Description Login to service using given credentials
// @Tags Guest user
// @Accept  json
// @Produce  json
// @Param User body userLoginRequest true "User credentials"
// @Success 200 {object} userResponse
// @Failure default {object} utils.Error
// @Router /users/login [post]
func (h *Handler) Login(c echo.Context) error {
	req := &userLoginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := h.userStore.GetByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	rsp, err := newUserResponse(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rsp)
}

// CurrentUser
// @Summary Get User by id in token
// @Description Get User by id in token
// @Tags users
// @Accept  json
// @Produce  json
// @Param User body userLoginRequest true "User credentials"
// @Success 200 {object} userResponse
// @Failure default {object} utils.Error
// @Router /users [get]
func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	rsp, err := newUserResponse(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rsp)
}

// UpdateUser
// @Summary Update user by id  in token
// @Description Update User by id in token. It can update
// @Tags users
// @Accept  json
// @Produce  json
// @Param User body userLoginRequest true "User credentials"
// @Success 200 {object} userResponse
// @Failure default {object} utils.Error
// @Router /users [put]
func (h *Handler) UpdateUser(c echo.Context) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := newUserUpdateRequest()
	req.populate(u)
	if err := req.bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.userStore.Update(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	rsp, err := newUserResponse(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rsp)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := newUserDeleteRequest()
	if err = req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = h.userStore.Delete(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	rsp, err := newUserResponse(u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rsp)
}

func userIDFromToken(c echo.Context) uuid.UUID {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(uuid.UUID)

	return id
}
