package hangman_test

import (
	"testing"

	"github.com/imhotepio/letsgo_labs/hangman"
	"github.com/stretchr/testify/assert"
)

func TestGameWon(t *testing.T) {
	g, err := hangman.NewGame("assets/test.1.txt")
	assert.Nil(t, err)

	uu := []struct {
		guess   rune
		status  hangman.GameStatus
		letters string
	}{
		{'t', hangman.CorrectGuess, "t--t"},
		{'t', hangman.AlreadyGuessed, "t--t"},
		{'z', hangman.IncorrectGuess, "t--t"},
		{'a', hangman.CorrectGuess, "t-at"},
		{'h', hangman.Won, "that"},
	}

	for _, u := range uu {
		ta, err := g.Guess(u.guess)

		assert.Nil(t, err)
		assert.Equal(t, u.status, ta.Status)
		assert.Equal(t, u.letters, string(ta.Letters))
	}
}

func TestGameLoose(t *testing.T) {
	g, err := hangman.NewGame("assets/test.1.txt")
	assert.Nil(t, err)

	uu := []struct {
		guess   rune
		turns   int
		status  hangman.GameStatus
		letters string
	}{
		{'b', 6, hangman.IncorrectGuess, "----"},
		{'c', 5, hangman.IncorrectGuess, "----"},
		{'d', 4, hangman.IncorrectGuess, "----"},
		{'e', 3, hangman.IncorrectGuess, "----"},
		{'f', 2, hangman.IncorrectGuess, "----"},
		{'g', 1, hangman.IncorrectGuess, "----"},
		{'i', 0, hangman.Lost, "----"},
	}

	for _, u := range uu {
		ta, err := g.Guess(u.guess)
		assert.Nil(t, err)
		assert.Equal(t, u.status, ta.Status)
		assert.Equal(t, u.turns, ta.TurnsLeft)
		assert.Equal(t, u.letters, string(ta.Letters))
	}
}

func TestGameAlreadyGuessed(t *testing.T) {
	g, err := hangman.NewGame("assets/test.1.txt")
	assert.Nil(t, err)

	uu := []struct {
		guess   rune
		turns   int
		status  hangman.GameStatus
		letters string
	}{
		{'b', 6, hangman.IncorrectGuess, "----"},
		{'b', 6, hangman.AlreadyGuessed, "----"},
	}

	for _, u := range uu {
		ta, err := g.Guess(u.guess)
		assert.Nil(t, err)
		assert.Equal(t, u.status, ta.Status)
		assert.Equal(t, u.turns, ta.TurnsLeft)
		assert.Equal(t, u.letters, string(ta.Letters))
	}
}
