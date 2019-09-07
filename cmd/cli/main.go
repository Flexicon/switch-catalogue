package main

import (
	"github.com/flexicon/switch-catalogue/pkg/cmd"
	"github.com/flexicon/switch-catalogue/pkg/db"
	"github.com/flexicon/switch-catalogue/pkg/fetching"
	"github.com/flexicon/switch-catalogue/pkg/listing"
	"github.com/flexicon/switch-catalogue/pkg/store"
	"github.com/flexicon/switch-catalogue/pkg/writing"
	"github.com/labstack/gommon/log"
)

func main() {
	d := db.New()

	gs := store.NewGameStore(d)
	gApi := fetching.NewGameApiService()

	lgs := listing.NewGameService(gs)
	wgs := writing.NewGameService(gs)

	c := cmd.New(lgs, wgs, gApi)
	c.RegisterFlags()
	c.RegisterCommands()

	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
