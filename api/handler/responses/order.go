package responses

import "github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"

type TokenResponse struct {
	Token string `param:"token" query:"token" form:"token" json:"token" xml:"token"`
}

func NewTokenResponse(token string) (*TokenResponse, error) {
	r := new(TokenResponse)
	r.Token = token
	return r, nil
}

type UserResponse struct {
	User struct {
		Username string         `param:"username" query:"username" form:"username" json:"username" xml:"username"`
		Email    string         `param:"email" query:"email" form:"email" json:"email" xml:"email"`
		Role     model.RoleName `param:"role" query:"role" form:"role" json:"role" xml:"role"`
	} `json:"user"`
}

func NewUserResponse(u *model.User) (*UserResponse, error) {
	r := new(UserResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Role = u.Role
	return r, nil
}
