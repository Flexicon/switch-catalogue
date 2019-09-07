package db

import (
	"fmt"
	"github.com/flexicon/switch-catalogue/pkg/store"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type config struct {
	host    string
	port    string
	user    string
	pass    string
	dbName  string
	verbose bool
}

func New() *gorm.DB {
	c := newConfig()
	db, err := gorm.Open("postgres", prepareConnectionString(c))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err))
	}

	if c.verbose {
		db.LogMode(true)
	}

	return db
}

func AutoMigrate(d *gorm.DB) {
	d.AutoMigrate(
		&store.Game{},
	)
}

func prepareConnectionString(c *config) string {
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

	verbose := false
	v := os.Getenv("DB_LOG_VERBOSE")
	if v == "true" {
		verbose = true
	}

	return &config{
		host:    host,
		port:    port,
		user:    user,
		pass:    pass,
		dbName:  dbName,
		verbose: verbose,
	}
}
