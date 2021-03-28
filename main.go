package main

import (
	"context"
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/dao/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
	"os/signal"
	"time"
)

type App struct {
	e *echo.Echo
}

func main() {
	app := App{}

	err := app.initialize()
	if err != nil {
		log.Fatalf("Could not initialize app: %v", err)
	}

	err = dbInitialize()
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Could not run app: %v", err)
	}
}

func (a *App) initialize() error {
	a.e = echo.New()

	// Middleware
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())

	return nil
}

func (a *App) Run() error {
	go func() {
		if err := a.e.Start(":1323"); err != nil {
			a.e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.e.Shutdown(ctx); err != nil {
		a.e.Logger.Fatal(err)
	}
	return nil
}

func dbInitialize() error {
	conn := postgres.Connection{}
	err := conn.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to postgres: %v", err)
	}
	err = conn.PrepareDB()

	return nil
}
