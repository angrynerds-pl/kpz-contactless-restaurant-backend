package api

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/auth"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao/postgres"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type CRUDGroup struct {
	G       *echo.Group
	Handler handler.CRUDHandler
}

func SetRouters(e *echo.Echo, db *gorm.DB) error {
	// Auth struct
	authS := auth.Auth{User: &postgres.User{Db: db}}

	e.POST("/login", authS.Login)

	//authGroup := e.Group("/auth")
	restrictedGroup := e.Group("/restricted")

	restrictedGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		//ErrorHandlerWithContext: authS.JWTErrorChecker,
		SigningKey: []byte(auth.JwtSecretKey),
	}))

	CRUDGroups := []CRUDGroup{
		{
			G:       restrictedGroup.Group("/users"),
			Handler: handler.NewCRUDHandler("user"),
		},
	}

	for _, CRUDGroup := range CRUDGroups {
		groupCRUDMethods(CRUDGroup.G, CRUDGroup.Handler)
	}

	//authGroupMethods(authGroup, &handler.Handler{})

	return nil
}

func groupCRUDMethods(g *echo.Group, h handler.CRUDHandler) {
	g.GET("", h.GetAll)
	g.POST("", h.Create)
	g.GET("/:id", h.Get)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

//func authGroupMethods(g *echo.Group, h *handler.Handler) {
//	g.POST("/login", h.Login)
//}
