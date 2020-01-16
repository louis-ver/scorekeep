package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddFavorite(t *testing.T) {
	colorado := "colorado-avalanche"
	montreal := "montreal-canadiens"

	config := Config{
		Favorites: Favorites{
			NHL: []string{colorado},
		},
	}

	config.AddFavorite(montreal, NHL)

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

	config.AddFavorite(colorado, NHL)

	assert.Equal(t, config.Favorites.NHL, []string{colorado})
}
