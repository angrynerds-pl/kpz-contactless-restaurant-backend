package requests

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/labstack/echo/v4"
)

type UserUpdateRequest struct {
	User struct {
		Username string `param:"username" query:"username" form:"username" json:"username" xml:"username validate:"required"`
		Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email" validate:"required,email"`
		Password string `param:"password" query:"password" form:"password" json:"password" xml:"password" validate:"required"`
	} `json:"user"`
}

func NewUserUpdateRequest() *UserUpdateRequest {
	return new(UserUpdateRequest)
}

func (r *UserUpdateRequest) Populate(u *model.User) {
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Password = u.Password
}

func (r *UserUpdateRequest) Bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	if r.User.Password != u.Password {
		h, err := u.HashPassword(r.User.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	return nil
}

type UserRegisterRequest struct {
	User struct {
		Username string `param:"username" query:"username" form:"username" json:"username" xml:"username validate:"required"`
		Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email" validate:"required,email"`
		Password string `param:"password" query:"password" form:"password" json:"password" xml:"password" validate:"required"`
	} `json:"user"`
}

func (r *UserRegisterRequest) Bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type UserLoginRequest struct {
	User struct {
		Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email" validate:"required,email"`
		Password string `param:"password" query:"password" form:"password" json:"password" xml:"password" validate:"required"`
	} `json:"user"`
}

func (r *UserLoginRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type UserDeleteRequest struct {
}

func NewUserDeleteRequest() *UserDeleteRequest {
	return &UserDeleteRequest{}
}

func (r *UserDeleteRequest) Bind(c echo.Context) error {
	return nil
}
