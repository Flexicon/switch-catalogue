package fetching

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	gameApiBaseUrl  = "http://search.nintendo-europe.com/en/select"
	paramFq         = "type:GAME AND system_type:nintendoswitch* AND product_code_txt:*"
	paramWt         = "json"
	paramQ          = "*"
	paramSortTitle  = "sorting_title asc"
	paramSortNewest = "change_date desc"
)

type GameApi interface {
	FetchGames(offset, limit int, newest bool) ([]*Game, error)
}

type GameApiService struct {
	httpClient *http.Client
}

func NewGameApiService() *GameApiService {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	return &GameApiService{
		httpClient: httpClient,
	}
}

func (s *GameApiService) FetchGames(offset, limit int, newest bool) ([]*Game, error) {
	var nresponse NResponse
	apiUrl := prepareUrl(offset, limit, newest)

	res, err := s.httpClient.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buf, &nresponse)
	if err != nil {
		return nil, err
	}

	return gamesListFromResponse(nresponse), nil
}

func prepareUrl(offset, limit int, newest bool) string {
	params := url.Values{}
	params.Add("start", fmt.Sprint(offset))
	params.Add("rows", fmt.Sprint(limit))
	params.Add("fq", paramFq)
	params.Add("wt", paramWt)
	params.Add("q", paramQ)

	if newest {
		params.Add("sort", paramSortNewest)
	} else {
		params.Add("sort", paramSortTitle)
	}

	apiUrl, _ := url.Parse(gameApiBaseUrl)

	apiUrl.RawQuery = params.Encode()

	return apiUrl.String()
}
