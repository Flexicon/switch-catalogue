package commandline

import (
	"github.com/urfave/cli"
)

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
