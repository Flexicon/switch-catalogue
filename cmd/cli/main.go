package main

import (
	"github.com/flexicon/switch-catalogue/pkg/adding"
	"github.com/flexicon/switch-catalogue/pkg/api"
	"github.com/flexicon/switch-catalogue/pkg/commandline"
	"github.com/flexicon/switch-catalogue/pkg/db"
	"github.com/flexicon/switch-catalogue/pkg/listing"
	"github.com/flexicon/switch-catalogue/pkg/store"
	"github.com/labstack/gommon/log"
)

func main() {
	d := db.New()
	db.AutoMigrate(d)

	gs := store.NewGameStore(d)
	gApi := api.NewGameApiService()

	lgs := listing.NewGameService(gs)
	ags := adding.NewGameService(gs)

	c := commandline.New(lgs, ags, gApi)
	c.RegisterFlags()
	c.RegisterCommands()

	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
