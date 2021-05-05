package menu

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	uuid "github.com/satori/go.uuid"
)

type Menu interface {
	AddFoodToMenu(menuId uuid.UUID, food *model.Food) error
	GetMenu(menuId uuid.UUID) (*model.Menu, error)
	RemoveFoodFromMenu(menuId uuid.UUID, food *model.Food) error
}
