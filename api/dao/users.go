package dao

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	uuid "github.com/satori/go.uuid"
)

type User interface {
	Create(user *model.User) (*model.User, error)
	Get(id string) (*model.User, error)
	GetByEmailAndPassword(email string, password string) (*model.User, error)
	//GetAll() ([]*model.User,  error)
	Delete(id uuid.UUID) (string, error)
	Update(user *model.User, userUpdate *model.User) (*model.User, error)
}
