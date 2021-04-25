package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Username string `param:"username" query:"username" form:"username" json:"username" xml:"username"`
	Email    string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
	Password string `param:"password" query:"password" form:"password" json:"password" xml:"password"`

	Role        RoleName     `param:"role" query:"role" form:"role" json:"role" xml:"role"`
	Restaurants []Restaurant `gorm:"foreignKey:OwnerId" param:"restaurants" query:"restaurants" form:"restaurants" json:"restaurants" xml:"restaurants" `
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

type RoleName int

const (
	Admin RoleName = iota
	Customer
	Owner
	Manager
	Employee
)

func ToRoleName(s float64) RoleName {
	return toID[s]
}

//func (s RoleName) String() string {
//	return toString[s]
//}
//
//var toString = map[RoleName]string{
//	Admin:  "Admin",
//	Customer:  "Customer",
//	Owner: "Owner",
//	Manager: "Manager",
//	Employee: "Employee",
//}

var toID = map[float64]RoleName{
	0: Admin,
	1: Customer,
	2: Owner,
	3: Manager,
	4: Employee,
}

//
//// MarshalJSON marshals the enum as a quoted json string
//func (s RoleName) MarshalJSON() ([]byte, error) {
//	buffer := bytes.NewBufferString(`"`)
//	buffer.WriteString(toString[s])
//	buffer.WriteString(`"`)
//	return buffer.Bytes(), nil
//}
//
//// UnmarshalJSON unmashals a quoted json string to the enum value
//func (s *RoleName) UnmarshalJSON(b []byte) error {
//	var j string
//	err := json.Unmarshal(b, &j)
//	if err != nil {
//		return err
//	}
//	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
//	*s = toID[j]
//	return nil
//}
