package main

import (
	"github.com/flexicon/switch-catalogue/pkg/db"
	"github.com/flexicon/switch-catalogue/pkg/http/handler"
	"github.com/flexicon/switch-catalogue/pkg/http/router"
	"github.com/flexicon/switch-catalogue/pkg/listing"
	"github.com/flexicon/switch-catalogue/pkg/store"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	setupEnvironment()

	r := router.New()
	base := r.Group("")
	api := r.Group("/api")

	d := db.New()
	db.AutoMigrate(d)

	gs := store.NewGameStore(d)

	lgs := listing.NewGameService(gs)

	h := handler.NewHandler(lgs)
	h.RegisterBase(base)
	h.RegisterApi(api)

	router.Run(r)
}
