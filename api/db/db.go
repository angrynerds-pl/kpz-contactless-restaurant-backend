package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func New() (*gorm.DB, error) {
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

	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else if debug == "false" {
		db.Logger = db.Logger.LogMode(logger.Silent)
	}
	if err != nil {
		return nil, err
	}

	return db, nil
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
		"public.users",
	)
}
