package commandline

import (
	"fmt"
	"github.com/urfave/cli"
)

func actionFetch(cmd *Cmd) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		isDryRun := c.Parent().Bool("dry-run")
		if isDryRun {
			printDryRunBanner()
		}
		printWithUnderline("Fetching games")

		apiGames, err := cmd.gameApi.FetchGames(0, 10)
		if err != nil {
			return err
		}

		for i, g := range apiGames {
			fmt.Printf("%2d %10s %s\n", i+1, g.FsId, g.Title)
		}

		fmt.Println("TODO: save new games to DB...")

		return nil
	}
}