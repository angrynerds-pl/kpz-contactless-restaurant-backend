package responses

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
)

type tokenResponse struct {
	Token string `param:"token" query:"token" form:"token" json:"token" xml:"token"`
}

func NewTokenResponse(token string) (*tokenResponse, error) {
	r := new(tokenResponse)
	r.Token = token
	return r, nil
}

type userResponse struct {
	User struct {
		Username string         `param:"username" query:"username" form:"username" json:"username" xml:"username"`
		Email    string         `param:"email" query:"email" form:"email" json:"email" xml:"email"`
		Role     model.RoleName `param:"role" query:"role" form:"role" json:"role" xml:"role"`
	} `json:"user"`
}

func NewUserResponse(u *model.User) (*userResponse, error) {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Role = u.Role
	return r, nil
}
