package cmd

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

func (cmd *Cmd) RegisterFlags() {
	cmd.app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dry-run, d",
			Usage: "Run a command without persisting changes",
		},
	}
}

func (cmd *Cmd) RegisterCommands() {
	cmd.app.Commands = []cli.Command{
		{
			Name:    "last",
			Aliases: []string{"l"},
			Usage:   "Show the games last added to the system",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "limit",
					Usage: "Limits the amount of games to show",
					Value: 10,
				},
			},
			Action: actionLast(cmd),
		},
		{
			Name:    "fetch",
			Aliases: []string{"f"},
			Usage:   "Scrape for new games",
			Action:  actionFetch(cmd),
		},
	}
}
