package api

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddFavorite(t *testing.T) {
	colorado := "colorado-avalanche"
	montreal := "montreal-canadiens"

	config := Config{
		Favorites: Favorites{
			NHL: []string{colorado},
		},
	}

	config.AddFavorite(montreal, nhl)

	assert.Equal(t, config.Favorites.NHL, []string{colorado, montreal})
}

func TestFavoriteNotInAnyLeague(t *testing.T) {
	montreal := "montreal-canadiens"
	config := Config{
		Favorites: Favorites{
			NHL: []string{montreal},
		},
	}

	config.AddFavorite(montreal, "not-supported-league")
}

func TestAddFavoriteAlreadyInFavorites(t *testing.T) {
	colorado := "colorado-avalanche"

	config := Config{
		Favorites: Favorites{
			NHL: []string{colorado},
		},
	}

	config.AddFavorite(colorado, nhl)

	assert.Equal(t, config.Favorites.NHL, []string{colorado})
}

type IOUtil struct {
	mock.Mock
}

func (m *IOUtil) ReadFile(filename string) ([]byte, error) {
	fmt.Println("Mocked ReadFile function")
	config, _ := yaml.Marshal(&Config{
		Favorites: Favorites{NHL: []string{"montreal-canadiens"}},
	})
	return config, nil
}

func TestGetConfig(t *testing.T) {
	ioUtil := new(IOUtil)
	filename := "/Users/louisolivierguerin/.scorekeep/config.yaml"
	expect := Config{
		Favorites: Favorites{NHL: []string{"montreal-canadiens"}},
	}
	byteConfig, _ := yaml.Marshal(&expect)

	ioUtil.On("ReadFile", filename).Return(byteConfig, nil)

	config := GetConfig()

	assert.Equal(t, expect, config)
}
