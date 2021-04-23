package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/utils"
)

type userResponse struct {
	Token string
	User  struct {
		Username string `param:"username" query:"username" form:"username" json:"username" xml:"username"`
		Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
		Role     string `param:"role" query:"role" form:"role" json:"role" xml:"role"`
		//Token    string `param:"token" query:"token" form:"token" json:"token" xml:"token" validate:"required"`
	} `json:"user"`
}

func newUserResponse(u *model.User) (*userResponse, error) {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	token, err := utils.GenerateJWT(u.ID, u.Role)
	if err != nil {
		return nil, err
	}
	//r.User.Token = token
	r.Token = token
	return r, nil
}
