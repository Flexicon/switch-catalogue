package fetching

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/flexicon/switch-catalogue/pkg/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) Get(url string) (resp *http.Response, err error) {
	args := m.Called(url)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*http.Response), args.Error(1)
}

func TestNewGameApiService(t *testing.T) {
	service := NewGameApiService()

	assert.IsType(t, &GameApiService{}, service)
}

func TestGameApiService_FetchGames_HttpError(t *testing.T) {
	httpError := errors.New("some http error")
	mockHC := &MockHttpClient{}
	mockHC.On("Get", mock.AnythingOfType("string")).Return(nil, httpError).Once()

	service := &GameApiService{httpClient: mockHC}

	games, err := service.FetchGames(0, 10, true)

	assert.Nil(t, games)
	if assert.Error(t, err) {
		assert.Equal(t, httpError, err)
	}
}

func TestGameApiService_FetchGames_JsonError(t *testing.T) {
	httpRes := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
	}
	mockHC := &MockHttpClient{}
	mockHC.On("Get", mock.AnythingOfType("string")).Return(httpRes, nil).Once()

	service := &GameApiService{httpClient: mockHC}

	games, err := service.FetchGames(0, 10, true)

	assert.Nil(t, games)
	assert.Error(t, err)
	assert.IsType(t, &json.SyntaxError{}, err)
}

func TestGameApiService_FetchGames_Success(t *testing.T) {
	httpRes := &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(sampleSingleGameResponse))),
	}
	mockHC := &MockHttpClient{}
	mockHC.On("Get", mock.AnythingOfType("string")).Return(httpRes, nil).Once()

	service := &GameApiService{httpClient: mockHC}

	games, err := service.FetchGames(0, 1, false)

	assert.Nil(t, err)
	assert.IsType(t, []*store.Game{}, games)
	assert.Len(t, games, 1)
}

const sampleSingleGameResponse = `
{
  "responseHeader": {
    "status": 0,
    "QTime": 0,
    "params": {
      "q": "*",
      "start": "0",
      "fq": "type:GAME AND system_type:nintendoswitch* AND product_code_txt:*",
      "rows": "1",
      "wt": "json"
    }
  },
  "response": {
    "numFound": 2498,
    "start": 0,
    "docs": [
      {
        "fs_id": "1633056",
        "change_date": "2019-09-04T10:31:17Z",
        "url": "/Games/Nintendo-Switch/Tokyo-Mirage-Sessions-FE-Encore-1633056.html",
        "type": "GAME",
        "dates_released_dts": [
          "2020-01-17T00:00:00Z"
        ],
        "club_nintendo": false,
        "pretty_date_s": "17/01/2020",
        "play_mode_tv_mode_b": true,
        "play_mode_handheld_mode_b": true,
        "product_code_txt": [
          "HACPASA4A"
        ],
        "image_url_sq_s": "//cdn01.nintendo-europe.com/media/images/11_square_images/games_18/nintendo_switch_5/SQ_NSwitch_TokyoMirageSessionsFEEncore_image500w.jpg",
        "pg_s": "GAME",
        "gift_finder_detail_page_image_url_s": "//cdn01.nintendo-europe.com/media/images/11_square_images/games_18/nintendo_switch_5/SQ_NSwitch_TokyoMirageSessionsFEEncore_gift_finder_detailpage.jpg",
        "compatible_controller": [
          "nintendoswitch_pro_controller"
        ],
        "image_url": "//cdn01.nintendo-europe.com/media/images/11_square_images/games_18/nintendo_switch_5/SQ_NSwitch_TokyoMirageSessionsFEEncore_image500w.jpg",
        "originally_for_t": "HAC",
        "paid_subscription_required_b": false,
        "cloud_saves_b": false,
        "priority": "2021-01-17T00:00:00Z",
        "digital_version_b": false,
        "title_extras_txt": [
          "Tokyo Mirage Sessions ♯FE Encore"
        ],
        "image_url_h2x1_s": "//cdn01.nintendo-europe.com/media/images/10_share_images/games_15/nintendo_switch_4/H2x1_NSwitch_TokyoMirageSessionsFEEncore_image500w.jpg",
        "system_type": [
          "nintendoswitch_gamecard,nintendoswitch_digitaldistribution"
        ],
        "age_rating_sorting_i": 12,
        "game_categories_txt": [
          "rpg"
        ],
        "play_mode_tabletop_mode_b": true,
        "publisher": "Nintendo",
        "product_code_ss": [
          "HACPASA4A"
        ],
        "excerpt": "The game that brought together the worlds of Fire Emblem and ATLUS is coming to Nintendo Switch.",
        "nsuid_txt": [
          "70010000023326"
        ],
        "date_from": "2020-01-17T00:00:00Z",
        "language_availability": [
          ""
        ],
        "price_has_discount_b": false,
        "price_discount_percentage_f": 0.0,
        "title": "Tokyo Mirage Sessions ♯FE Encore",
        "sorting_title": "Tokyo Mirage Sessions ♯FE Encore",
        "copyright_s": "©2015-2020 Nintendo / ATLUS\nFIRE EMBLEM SERIES / SÉRIE FIRE EMBLEM : ©Nintendo / INTELLIGENT SYSTEMS",
        "gift_finder_carousel_image_url_s": "//cdn01.nintendo-europe.com/media/images/11_square_images/games_18/nintendo_switch_5/SQ_NSwitch_TokyoMirageSessionsFEEncore_gift_finder_carousel.jpg",
        "wishlist_email_square_image_url_s": "//cdn01.nintendo-europe.com/media/images/11_square_images/games_18/nintendo_switch_5/SQ_NSwitch_TokyoMirageSessionsFEEncore_square_image_wishlist.jpg",
        "players_to": 1,
        "wishlist_email_banner640w_image_url_s": "//cdn01.nintendo-europe.com/media/images/10_share_images/games_15/nintendo_switch_4/H2x1_NSwitch_TokyoMirageSessionsFEEncore_banner_image_wishlist_640w.jpg",
        "voice_chat_b": false,
        "playable_on_txt": [
          "HAC"
        ],
        "hits_i": 109,
        "pretty_game_categories_txt": [
          "RPG"
        ],
        "gift_finder_wishlist_image_url_s": "//cdn01.nintendo-europe.com/media/images/11_square_images/games_18/nintendo_switch_5/SQ_NSwitch_TokyoMirageSessionsFEEncore_gift_finder_wishlist.jpg",
        "switch_game_voucher_b": false,
        "game_category": [
          "rpg"
        ],
        "system_names_txt": [
          "Nintendo Switch"
        ],
        "pretty_agerating_s": "12",
        "players_from": 1,
        "age_rating_type": "pegi",
        "price_sorting_f": 49.99,
        "price_lowest_f": 49.99,
        "age_rating_value": "12",
        "physical_version_b": true,
        "wishlist_email_banner460w_image_url_s": "//cdn01.nintendo-europe.com/media/images/10_share_images/games_15/nintendo_switch_4/H2x1_NSwitch_TokyoMirageSessionsFEEncore_banner_image_wishlist_460w.jpg",
        "_version_": 1643985404642197507
      }
    ]
  }
}
`
