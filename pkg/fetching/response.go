package fetching

type NGameResponse struct {
	Title          string   `json:"title"`
	ProductCodeTxt []string `json:"product_code_txt"`
	FsId           string   `json:"fs_id"`
	Url            string   `json:"url"`
}

type NResponse struct {
	Response struct {
		Docs []NGameResponse `json:"docs"`
	} `json:"response"`
}

type Game struct {
	Title       string
	ProductCode string
	FsId        string
	Url         string
}

func newGameFromResponse(r NGameResponse) *Game {
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
