package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	gameApiBaseUrl = "http://search.nintendo-europe.com/en/select?fq=type%3AGAME%20AND%20system_type%3Anintendoswitch*%20AND%20product_code_txt%3A*&q=*&start=0&wt=json&rows=20"
)

type GameApi interface {
	FetchGames(offset, limit int) ([]*Game, error)
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

func (s *GameApiService) FetchGames(offset, limit int) ([]*Game, error) {
	var nresponse NResponse

	res, err := s.httpClient.Get(gameApiBaseUrl)
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
