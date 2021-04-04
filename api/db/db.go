package db

import (
	"github.com/angrynerds-pl/kpz-contactless-restaurant-backend/api/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func New() *gorm.DB {
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
	return db
}

//func TestDB() *gorm.DB {
//	db, err := gorm.Open("sqlite3", "./../realworld_test.db")
//	if err != nil {
//		fmt.Println("storage err: ", err)
//	}
//	db.DB().SetMaxIdleConns(3)
//	db.LogMode(false)
//	return db
//}
//
//func DropTestDB() error {
//	if err := os.Remove("./../realworld_test.db"); err != nil {
//		return err
//	}
//	return nil
//}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
	)
}
