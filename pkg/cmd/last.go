package cmd

import (
	"fmt"
	"github.com/urfave/cli"
)

func actionLast(cmd *Cmd) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		printWithUnderline("Last added games")

		limit := c.Int("limit")
		games, count, err := cmd.listingGameService.LastAdded(limit)
		if err != nil {
			return err
		}

		fmt.Printf("Total games in system: %d\n\n", count)
		fmt.Println("Last added:")

		for _, g := range games {
			fmt.Printf("ID: %5d - %s\n", g.ID, g.Title)
		}

		return nil
	}
}
