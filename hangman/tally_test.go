package hangman_test

import (
	"testing"

	"github.com/imhotep/letsgo_labs/hangman"
	"github.com/stretchr/testify/assert"
)

func TestNewTally(t *testing.T) {
	ta := hangman.NewTally(7, 10)

	assert.Equal(t, 7, ta.TurnsLeft)
}

func TestUpdateBadGuess(t *testing.T) {
	var (
		w        = "bumblebeetuna"
		ta       = newTally(w)
		expected = "_____________"
	)
	ta.Update(w, 'z')

	assert.Equal(t, 6, ta.TurnsLeft)
	assert.Equal(t, hangman.Active, ta.Status)
	assert.Equal(t, []rune(expected), ta.Letters)
}

func TestUpdateGoodGuess(t *testing.T) {
	var (
		w        = "bumblebeetuna"
		ta       = newTally(w)
		expected = "_____e_ee____"
	)
	ta.Update(w, 'e')

	assert.Equal(t, 7, ta.TurnsLeft)
	assert.Equal(t, hangman.Active, ta.Status)
	assert.Equal(t, []rune(expected), ta.Letters)
}

func TestUpdateAlreadyGuessed(t *testing.T) {
	var (
		w        = "bumblebeetuna"
		ta       = newTally(w)
		expected = "_____e_ee____"
	)
	ta.Update(w, 'e')
	ta.Update(w, 'e')

	assert.Equal(t, 7, ta.TurnsLeft)
	assert.Equal(t, hangman.AlreadyGuessed, ta.Status)
	assert.Equal(t, []rune(expected), ta.Letters)
}

func TestUpdateTallyWon(t *testing.T) {
	var (
		w  = "test"
		ta = newTally(w)
	)

	for _, l := range []rune("tes") {
		ta.Update(w, l)
	}

	assert.Equal(t, 7, ta.TurnsLeft)
	assert.Equal(t, hangman.Won, ta.Status)
	assert.Equal(t, []rune(w), ta.Letters)
}

func TestUpdateTallyLost(t *testing.T) {
	var (
		w  = "test"
		ta = newTally(w)
	)

	badGuesses := "abcdfgh"
	for _, c := range badGuesses {
		ta.Update(w, c)
	}

	assert.Equal(t, 0, ta.TurnsLeft)
	assert.Equal(t, hangman.Lost, ta.Status)
	assert.Equal(t, []rune("____"), ta.Letters)
}

func newTally(w string) *hangman.Tally {
	return hangman.NewTally(7, len(w))
}
