package store

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MenuStore struct {
	db *gorm.DB
}

func NewMenuStore(db *gorm.DB) *MenuStore {
	return &MenuStore{
		db: db,
	}
}

func (ms MenuStore) AddFoodToMenu(menu *model.Menu, food *model.Food) error {
	var m model.Menu
	m.ID = menu.ID

	err := ms.db.Model(&m).Association("Foods").Append(food)
	if err != nil {
		return err
	}
	return nil
}

func (ms MenuStore) GetMenu(menuId uuid.UUID) (*model.Menu, error) {
	var m model.Menu
	m.ID = menuId

	err := ms.db.Preload(clause.Associations).First(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (ms MenuStore) RemoveFoodFromMenu(menuId uuid.UUID, foodId uuid.UUID) error {
	var m model.Menu
	m.ID = menuId

	var f model.Food
	f.ID = foodId

	err := ms.db.Model(&m).Association("Foods").Delete(f)
	if err != nil {
		return err
	}
	return nil
}
