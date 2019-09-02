package fetching

type GameResponse struct {
	Title          string   `json:"title"`
	ProductCodeTxt []string `json:"product_code_txt"`
	FsId           string   `json:"fs_id"`
	Url            string   `json:"url"`
}

type NResponse struct {
	Response struct {
		Docs []GameResponse `json:"docs"`
	} `json:"response"`
}
