package store

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RestaurantStore struct {
	db *gorm.DB
}

func NewRestaurantStore(db *gorm.DB) *RestaurantStore {
	return &RestaurantStore{
		db: db,
	}
}

func (rs RestaurantStore) AddRestaurantToUser(userId uuid.UUID, restaurant *model.Restaurant) error {
	var u model.User
	u.ID = userId

	if err := rs.db.Model(&u).Association("Restaurants").Append(restaurant); err != nil {
		return err
	}

	err := rs.db.Model(restaurant).Association("Address").Append(&restaurant.Address)
	if err != nil {
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

func (rs RestaurantStore) Update(userId uuid.UUID, restaurant *model.Restaurant) (*model.Restaurant, error) {
	restToUpdate := &model.Restaurant{}
	//restToUpdate.ID = restaurant.ID

	u := &model.User{}
	u.ID = userId

	err := rs.db.Preload(clause.Associations).Model(u).Where("id = ?", restaurant.ID.String()).
		Association("Restaurants").Find(&u.Restaurants)

	err = rs.db.Preload("Restaurants").Preload("Restaurants.Address").Model(&u.Restaurants[0]).Updates(restaurant).Error
	if err != nil {
		return nil, err
	}

	addr := &model.Address{}
	addr.RestaurantID = restaurant.ID
	err = rs.db.Model(addr).Where("restaurant_id = ?", restaurant.ID.String()).Updates(restaurant.Address).Error
	if err != nil {
		return nil, err
	}

	err = rs.db.Preload("Address").First(restToUpdate, "id = ?", u.Restaurants[0].ID).Error
	if err != nil {
		return nil, err
	}
	return restToUpdate, nil
}

func (rs RestaurantStore) DeleteRestaurantFromUser(userId, restaurantId uuid.UUID) error {
	user := &model.User{}
	user.ID = userId

	restaurant := &model.Restaurant{}
	restaurant.ID = restaurantId

	err := rs.db.Model(user).Association("Restaurants").Delete(restaurant)
	if err != nil {
		return err
	}
	return nil
}

func (rs RestaurantStore) AddAddressToRestaurant(restaurantId uuid.UUID, addr *model.Address) error {
	r := &model.Restaurant{}
	r.ID = restaurantId

	err := rs.db.Model(r).Association("Address").Replace(addr)
	if err != nil {
		return err
	}
	return nil
}
