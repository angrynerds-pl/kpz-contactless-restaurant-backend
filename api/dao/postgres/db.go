package postgres

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Connection struct {
	Db *gorm.DB
}

//ConnectDB inits the database connection
func (c *Connection) ConnectDB() error {
	err := godotenv.Load()

	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	},
	),
		&gorm.Config{
			PrepareStmt: true,
		},
	)
	if err != nil {
		panic("DB connection Error")
	}
	c.Db = db
	return nil
}

func (c *Connection) PrepareDB() error {
	err := c.Db.AutoMigrate(
		&model.User{},
		&model.Role{},
	)
	if err != nil {
		return err
	}
	return err
}
