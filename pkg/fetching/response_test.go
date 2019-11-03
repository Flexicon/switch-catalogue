package fetching

import (
	"github.com/flexicon/switch-catalogue/pkg/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGameResponse(t *testing.T) {
	actual := newGameFromResponse(nGameResponse{
		Title:          "Super Mario",
		ProductCodeTxt: []string{"S12345"},
		FsId:           "FS_123",
		Url:            "https://some.url.com",
	})
	expected := &store.Game{
		Title:       "Super Mario",
		ProductCode: "S12345",
		FsId:        "FS_123",
		Url:         "https://some.url.com",
	}

	assert.Equal(t, expected, actual)
}

func TestGamesListFromResponse(t *testing.T) {
	testResponse := nResponse{
		Response: nNestedResponse{Docs: []nGameResponse{
			{
				Title:          "Super Mario",
				ProductCodeTxt: []string{"S12345"},
				FsId:           "FS_123",
				Url:            "https://some.url.com",
			},
			{
				Title:          "Legend of Zelda",
				ProductCodeTxt: []string{"Q98765"},
				FsId:           "FS_345",
				Url:            "https://some.other.com/zelda",
			},
		}},
	}
	actual := gamesListFromResponse(testResponse)
	expected := []*store.Game{
		{
			Title:       "Super Mario",
			ProductCode: "S12345",
			FsId:        "FS_123",
			Url:         "https://some.url.com",
		},
		{
			Title:       "Legend of Zelda",
			ProductCode: "Q98765",
			FsId:        "FS_345",
			Url:         "https://some.other.com/zelda",
		},
	}

	assert.Equal(t, expected, actual)
}
