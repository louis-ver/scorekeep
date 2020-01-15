package main

import (
	"testing"

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

var byteArray = []byte{
	102, 97, 118, 111, 114, 105,
	116, 101, 115, 58, 10, 32, 32,
	110, 104, 108, 58, 10, 32, 45,
	32, 109, 111, 110, 116, 101, 97,
	108, 45, 99, 97, 100, 105, 101, 110, 115, 10}

func (m *IOUtil) ReadFile(filename string) ([]byte, error) {
	return byteArray, nil
}

func TestGetConfig(t *testing.T) {
	ioUtilMock := new(IOUtil)
	filename := "/.scorekeep/config.yaml"
	expectedConfig := Config{
		Favorites: Favorites{NHL: []string{"montreal-canadiens"}},
	}

	ioUtilMock.On("ReadFile", filename).Return(byteArray, nil)

	config := GetConfig()

	assert.Equal(t, expectedConfig, config)
}
