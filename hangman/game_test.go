package hangman_test

import (
	"testing"

	"github.com/imhotep/letsgo_labs/hangman"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := hangman.NewGame("test")

	assert.Equal(t, 7, g.Tally.TurnsLeft)
	assert.Equal(t, hangman.Active, g.Tally.Status)
}

func TestWin(t *testing.T) {
	g := hangman.NewGame("test")

	for _, l := range "tes" {
		g.Guess(l)
	}

	assert.Equal(t, 7, g.Tally.TurnsLeft)
	assert.Equal(t, hangman.Won, g.Tally.Status)
}

func TestLoose(t *testing.T) {
	g := hangman.NewGame("test")

	for _, l := range "abcdfghi" {
		g.Guess(l)
	}

	assert.Equal(t, 0, g.Tally.TurnsLeft)
	assert.Equal(t, hangman.Lost, g.Tally.Status)
}

func TestAlreadyGuessed(t *testing.T) {
	g := hangman.NewGame("test")

	for _, l := range "ttt" {
		g.Guess(l)
	}

	assert.Equal(t, 7, g.Tally.TurnsLeft)
	assert.Equal(t, hangman.AlreadyGuessed, g.Tally.Status)
}
