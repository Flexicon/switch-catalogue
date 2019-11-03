package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

func actionFetch(cmd *Cmd) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		isDryRun := c.Parent().Bool("dry-run")
		if isDryRun {
			fmt.Println(dryRunBanner())
		}
		fmt.Println(messageWithUnderline("Fetching games from external API"))

		batch := 0
		gameIndex := 0
		for {
			games, err := cmd.gameApi.FetchGames(batch*1000, 1000, true)
			batch += 1
			if err != nil {
				return err
			}

			for _, g := range games {
				fmt.Printf("%4d %10s %s\n", gameIndex+1, g.FsId, g.Title)
				gameIndex += 1
			}

			if !isDryRun {
				// Don't actually save the games in a dry run
				err = cmd.writingGameService.BatchUpsert(games)
				if err != nil {
					return err
				}
			}

			if len(games) < 1000 {
				return nil
			}
		}
	}
}
