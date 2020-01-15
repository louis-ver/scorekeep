package main

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

	config.AddFavorite(montreal, nhl)

	assert.Equal(t, config.Favorites.NHL, []string{colorado, montreal})
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
