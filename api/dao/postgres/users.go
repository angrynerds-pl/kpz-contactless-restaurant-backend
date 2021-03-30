package postgres

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	Db *gorm.DB
}

func (u *User) Create(user *model.User) (*model.User, error) {
	err := u.Db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Get(id string) (*model.User, error) {
	user := &model.User{}

	err := u.Db.First(user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) GetByEmailAndPassword(email string, password string) (*model.User, error) {
	user := &model.User{}
	err := u.Db.First(user, "email = ? AND password = ?", email, password).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

//func (u *User) GetAll() ([]*model.User, error) {
//	var users []model.User
//	user := &model.User{}
//
//	err := u.Db.Where("")
//
//}

func (u *User) Delete(id uuid.UUID) (string, error) {
	err := u.Db.Delete(&model.User{}, "id = ?", id.String()).Error
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (u *User) Update(user *model.User, userUpdate *model.User) (*model.User, error) {
	err := u.Db.Model(user).Updates(userUpdate).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
