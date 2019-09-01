package api

type Game struct {
	Title       string
	ProductCode string
	FsId        string
	Url         string
}

func newGameFromResponse(r GameResponse) *Game {
	return &Game{
		Title:       r.Title,
		ProductCode: r.ProductCodeTxt[0],
		FsId:        r.FsId,
		Url:         r.Url,
	}
}

func gamesListFromResponse(r NResponse) []*Game {
	games := make([]*Game, 0)

	for _, g := range r.Response.Docs {
		games = append(games, newGameFromResponse(g))
	}

	return games
}
