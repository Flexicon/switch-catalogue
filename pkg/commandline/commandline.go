package commandline

import (
	"github.com/flexicon/switch-catalogue/pkg/fetching"
	"github.com/flexicon/switch-catalogue/pkg/game"
	"github.com/urfave/cli"
	"os"
)

type Cmd struct {
	app                *cli.App
	listingGameService game.ListingService
	addingGameService  game.AddingService
	gameApi            fetching.GameApi
}

func New(lgs game.ListingService, ags game.AddingService, ga fetching.GameApi) *Cmd {
	app := cli.NewApp()
	app.Name = "switch-catalogue cli"
	app.Usage = "Manage the switch catalogue from the command line"
	app.Author = "Michal Repec (Flexicon)"
	app.Version = "1.0.0"

	return &Cmd{
		app:                app,
		listingGameService: lgs,
		addingGameService:  ags,
		gameApi:            ga,
	}
}

func (cmd *Cmd) Run() error {
	return cmd.app.Run(os.Args)
}
