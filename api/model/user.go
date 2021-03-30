package model

import uuid "github.com/satori/go.uuid"

type User struct {
	Base

	Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
	Password string `param:"password" query:"password" form:"password" json:"password" xml:"password"`
	Role     *Role
}

type Role struct {
	Base
	UserId uuid.UUID
}
