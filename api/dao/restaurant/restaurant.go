package restaurant

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	uuid "github.com/satori/go.uuid"
)

type Store interface {
	GetByID(userId, restaurantId uuid.UUID) (*model.Restaurant, error)
	GetAll(userId uuid.UUID) ([]model.Restaurant, error)
	AddRestaurantToUser(userId uuid.UUID, restaurant model.Restaurant) error
	Update(*model.Restaurant) error
	Delete(user *model.Restaurant) error
}
