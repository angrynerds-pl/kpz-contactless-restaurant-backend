package handler

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	guestUsers := v1.Group("/auth")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/customer", h.SignUpCustomer)
	guestUsers.POST("/owner", h.SignUpOwner)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/users", jwtMiddleware)
	user.GET("", h.CurrentUser)
	user.PUT("", h.UpdateUser)

	user.DELETE("", h.DeleteUser)

	{
		userRestaurants := user.Group("/restaurants", jwtMiddleware)
		userRestaurants.POST("", h.CreateRestaurant)
		userRestaurants.GET("/:id", h.GetRestaurant)
		userRestaurants.GET("", h.Restaurants)
		userRestaurants.PUT("/:id", h.UpdateRestaurant)
		userRestaurants.DELETE("/:id", h.RemoveRestaurantFromUser)
		userRestaurants.PUT("/:id/address", h.AddAddressToRestaurant)
	}

	// /users/restaurants/menus
	foodRestaurants := userRestaurants.Group("/menus", jwtMiddleware)
	foodRestaurants.POST("/:menuId ", h.AddFoodToMenu)
	foodRestaurants.GET("/:id", h.GetMenu)
	foodRestaurants.DELETE("/:id", h.RemoveFoodFromMenu)

}
