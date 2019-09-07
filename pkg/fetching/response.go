package fetching

import "time"

type nGameResponse struct {
	Title          string    `json:"title"`
	ProductCodeTxt []string  `json:"product_code_txt"`
	FsId           string    `json:"fs_id"`
	Url            string    `json:"url"`
	ChangeDate     time.Time `json:"change_date"`
}

type nNestedResponse struct {
	Docs []nGameResponse `json:"docs"`
}

type nResponse struct {
	Response nNestedResponse `json:"response"`
}

type Game struct {
	Title       string
	ProductCode string
	FsId        string
	Url         string
	ChangeDate  time.Time
}

func newGameFromResponse(r nGameResponse) *Game {
	return &Game{
		Title:       r.Title,
		ProductCode: r.ProductCodeTxt[0],
		FsId:        r.FsId,
		Url:         r.Url,
		ChangeDate:  r.ChangeDate,
	}
}

func gamesListFromResponse(r nResponse) []*Game {
	games := make([]*Game, 0)

	for _, g := range r.Response.Docs {
		games = append(games, newGameFromResponse(g))
	}

	return games
}
