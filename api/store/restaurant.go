package store

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type RestaurantStore struct {
	db *gorm.DB
}

func NewRestaurantStore(db *gorm.DB) *RestaurantStore {
	return &RestaurantStore{
		db: db,
	}
}

func (rs RestaurantStore) AddRestaurantToUser(userId uuid.UUID, restaurant model.Restaurant) error {
	var u model.User
	u.ID = userId

	if err := rs.db.Model(u).Association("Restaurants").Append(restaurant); err != nil {
		return err
	}
	return nil
}

func (rs RestaurantStore) GetByID(userId, restaurantId uuid.UUID) (*model.Restaurant, error) {
	u := model.User{}
	u.ID = userId

	err := rs.db.Model(u).Where("id = ?", restaurantId.String()).Association("Restaurants").Find(&u.Restaurants)
	if err != nil {
		return nil, err
	}
	return &u.Restaurants[0], nil
}

func (rs RestaurantStore) GetAll(userId uuid.UUID) ([]model.Restaurant, error) {
	u := model.User{}
	u.ID = userId

	err := rs.db.Model(u).Association("Restaurants").Find(&u.Restaurants)
	if err != nil {
		return nil, err
	}
	return u.Restaurants, nil
}

func (rs RestaurantStore) Update(restaurant *model.Restaurant) error {
	panic("implement me")
}

func (rs RestaurantStore) Delete(user *model.Restaurant) error {
	panic("implement me")
}
