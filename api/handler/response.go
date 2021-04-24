package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
)

type userResponse struct {
	User struct {
		Username string `param:"username" query:"username" form:"username" json:"username" xml:"username"`
		Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
		Role     string `param:"role" query:"role" form:"role" json:"role" xml:"role"`
	} `json:"user"`
}

func newUserResponse(u *model.User) (*userResponse, error) {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	return r, nil
}
