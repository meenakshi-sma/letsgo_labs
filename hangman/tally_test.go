package hangman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTally(t *testing.T) {
	ta := newTally(7, "blee")

	assert.Equal(t, 7, ta.TurnsLeft)
}

func TestUpdateWin(t *testing.T) {
	ta := newTally(7, "blee")
	for _, r := range []rune{'b', 'l'} {
		ta.update(r)
		assert.Equal(t, CorrectGuess, ta.Status)
	}
	ta.update('e')
	assert.Equal(t, Won, ta.Status)
}

func TestUpdateLoose(t *testing.T) {
	ta := newTally(3, "blee")
	for _, r := range []rune{'x', 'y'} {
		ta.update(r)
		assert.Equal(t, IncorrectGuess, ta.Status)
	}
	ta.update('z')
	assert.Equal(t, Lost, ta.Status)
}

func TestUpdateAlreadyGuessed(t *testing.T) {
	ta := newTally(3, "blee")
	ta.update('b')
	ta.update('b')
	assert.Equal(t, AlreadyGuessed, ta.Status)
}
