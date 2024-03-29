package fetching

import (
	"encoding/json"
	"fmt"
	"github.com/flexicon/switch-catalogue/pkg/store"
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
	FetchGames(offset, limit int, newest bool) ([]*store.Game, error)
}

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type GameApiService struct {
	httpClient httpClient
}

func NewGameApiService() *GameApiService {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	return &GameApiService{
		httpClient: httpClient,
	}
}

func (s *GameApiService) FetchGames(offset, limit int, newest bool) ([]*store.Game, error) {
	var nResponse nResponse
	apiUrl := prepareUrl(offset, limit, newest)

	res, err := s.httpClient.Get(apiUrl)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(buf, &nResponse)
	if err != nil {
		return nil, err
	}

	return gamesListFromResponse(nResponse), nil
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
