package db

import (
	"fmt"
	"github.com/flexicon/switch-catalogue/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func New() *gorm.DB {
	db, err := gorm.Open("postgres", prepareConnectionString())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err))
	}

	db.LogMode(true)

	return db
}

func prepareConnectionString() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "api"
	}

	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "pass"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "switch-catalogue"
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbName,
		pass,
	)
}

func AutoMigrate(d *gorm.DB) {
	d.AutoMigrate(
		&model.Game{},
	)
}
