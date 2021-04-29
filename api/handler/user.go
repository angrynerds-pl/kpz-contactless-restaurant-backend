package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler/requests"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler/responses"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Signup
// @Summary Register new user with owner role
// @Description Register new user in server
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param User body requests.UserRegisterRequest true "User credentials"
// @Success 200 {object} responses.UserResponse
// @Failure default {object} utils.Error
// @Router /users [post]
// @Deprecated
func (h *Handler) SignUp(c echo.Context) error {
	var u model.User
	req := &requests.UserRegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	token, err := utils.GenerateJWT(u.ID, u.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	rsp, err := responses.NewTokenResponse(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, rsp)
}

// Signup
// @Summary Register new user with customer role
// @Description Register new user in server
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param User body requests.UserRegisterRequest true "User credentials"
// @Success 200 {object} responses.TokenResponse
// @Failure default {object} utils.Error
// @Router /auth/customer [post]
func (h *Handler) SignUpCustomer(c echo.Context) error {
	var u model.User
	req := &requests.UserRegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u.Role = model.Customer
	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	token, err := utils.GenerateJWT(u.ID, u.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	rsp, err := responses.NewTokenResponse(token)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, rsp)
}

// Signup
// @Summary Register new user with owner role
// @Description Register new user in server
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param User body requests.UserRegisterRequest true "User credentials"
// @Success 200 {object} responses.TokenResponse
// @Failure default {object} utils.Error
// @Router /auth/owner [post]
func (h *Handler) SignUpOwner(c echo.Context) error {
	var u model.User
	req := &requests.UserRegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u.Role = model.Owner

	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	token, err := utils.GenerateJWT(u.ID, u.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	rsp, err := responses.NewTokenResponse(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, rsp)
}

// Login
// @Summary Login to service
// @Description Login to service using given credentials
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param User body requests.UserLoginRequest true "User credentials"
// @Success 200 {object} responses.TokenResponse
// @Failure default {object} utils.Error
// @Router /auth/login [post]
func (h *Handler) Login(c echo.Context) error {
	req := &requests.UserLoginRequest{}
	if err := req.Bind(c); err != nil {
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

	token, err := utils.GenerateJWT(u.ID, u.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	rsp, err := responses.NewTokenResponse(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, rsp)
}

// CurrentUser
// @Summary Get User by id in token
// @Description Get User by id in token
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.UserResponse
// @Failure default {object} utils.Error
// @Router /users [get]
// @Security Bearer
func (h *Handler) CurrentUser(c echo.Context) error {
	userId, err := userIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	u, err := h.userStore.GetByID(*userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	rsp, err := responses.NewUserResponse(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, rsp)
}

// UpdateUser
// @Summary Update user by id  in token
// @Description Update User by id in token. It can update
// @Tags Users
// @Accept  json
// @Produce  json
// @Param User body requests.UserLoginRequest true "User credentials"
// @Success 200 {object} responses.UserResponse
// @Failure default {object} utils.Error
// @Router /users [put]
func (h *Handler) UpdateUser(c echo.Context) error {
	userId, err := userIDFromToken(c)
	if err != nil {
		return err
	}
	u, err := h.userStore.GetByID(*userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := requests.NewUserUpdateRequest()
	req.Populate(u)
	if err := req.Bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.userStore.Update(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	rsp, err := responses.NewUserResponse(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, rsp)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	userId, err := userIDFromToken(c)
	if err != nil {
		return err
	}
	u, err := h.userStore.GetByID(*userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := requests.NewUserDeleteRequest()
	if err = req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = h.userStore.Delete(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	rsp, err := responses.NewUserResponse(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, rsp)
}
