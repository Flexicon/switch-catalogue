package db

import (
	"fmt"
	"github.com/flexicon/switch-catalogue/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type config struct {
	host   string
	port   string
	user   string
	pass   string
	dbName string
}

func New() *gorm.DB {
	db, err := gorm.Open("postgres", prepareConnectionString())
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err))
	}

	db.LogMode(true)

	return db
}

func AutoMigrate(d *gorm.DB) {
	d.AutoMigrate(
		&model.Game{},
	)
}

func prepareConnectionString() string {
	c := newConfig()

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.host,
		c.port,
		c.user,
		c.dbName,
		c.pass,
	)
}

func newConfig() *config {
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

	return &config{
		host:   host,
		port:   port,
		user:   user,
		pass:   pass,
		dbName: dbName,
	}
}
