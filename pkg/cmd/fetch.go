package cmd

import (
	"fmt"
	"github.com/flexicon/switch-catalogue/pkg/fetching"
	"github.com/flexicon/switch-catalogue/pkg/store"
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
			apiGames, err := cmd.gameApi.FetchGames(batch*1000, 1000, true)
			batch += 1
			if err != nil {
				return err
			}

			for _, g := range apiGames {
				fmt.Printf("%4d %10s %s\n", gameIndex+1, g.FsId, g.Title)
				gameIndex += 1
			}

			games := convertApiGamesToModels(apiGames)

			err = cmd.writingGameService.BatchUpsert(games)
			if err != nil {
				return err
			}

			if len(games) < 1000 {
				return nil
			}
		}
	}
}

// TODO: move this logic to the fetching package. Using the store.Game struct in more places should simplify a lot
func convertApiGamesToModels(games []*fetching.Game) []*store.Game {
	models := make([]*store.Game, 0)

	for _, g := range games {
		model := &store.Game{
			Title:       g.Title,
			ProductCode: g.ProductCode,
			FsId:        g.FsId,
			Url:         g.Url,
			ChangeDate:  g.ChangeDate,
		}

		models = append(models, model)
	}

	return models
}
