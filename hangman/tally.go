package hangman

import (
	"math/rand"
	"time"
)

const (
	// Started denotes a game in progres
	Started GameStatus = iota
	// Won denotes the game is won
	Won
	// Lost denotes the game is lost
	Lost
	// AlreadyGuessed denotes a letter was already given
	AlreadyGuessed
	// IncorrectGuess reports the letter was incorrect
	IncorrectGuess
	// CorrectGuess reports the letter was correct
	CorrectGuess
)

type (
	// GameStatus tracks the status of a game
	GameStatus int

	// Tally tracks the player status
	Tally struct {
		TurnsLeft int
		Status    GameStatus
		Letters   []rune
		guesses   []rune
		Word      string
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewTally initializes a new tally with a given number of turns
func newTally(maxTurns int, word string) *Tally {
	return &Tally{
		TurnsLeft: maxTurns,
		Word:      word,
		Letters:   lettersFromGuess(word, []rune{}),
	}
}

// Update the current tally
func (t *Tally) update(g rune) {
	if inGuesses(t.guesses, g) {
		t.Status = AlreadyGuessed
		return
	}

	t.guesses = append(t.guesses, g)
	t.Letters = lettersFromGuess(t.Word, t.guesses)

	if inWord(t.Word, g) {
		t.Status = CorrectGuess
		if !missingLetters(t.Letters) {
			t.Status = Won
		}
		return
	}

	t.TurnsLeft--
	t.Status = IncorrectGuess
	if t.TurnsLeft == 0 {
		t.Status = Lost
	}
}

func (t *Tally) reset(guesses int, word string) {
	t.Status = Started
	t.TurnsLeft = guesses
	t.Letters = lettersFromGuess(word, []rune{})
	t.Word = word
}

func (t *Tally) won() bool {
	return t.TurnsLeft > 0 && !missingLetters(t.Letters)
}

func missingLetters(ll []rune) bool {
	for _, l := range ll {
		if l == '-' {
			return true
		}
	}
	return false
}

func inWord(w string, g rune) bool {
	for _, r := range w {
		if r == g {
			return true
		}
	}
	return false
}

func lettersFromGuess(w string, guesses []rune) []rune {
	res := make([]rune, len(w))
	for i, r := range w {
		if inGuesses(guesses, r) {
			res[i] = r
		} else {
			res[i] = '-'
		}
	}
	return res
}

func inGuesses(guesses []rune, g rune) bool {
	for _, r := range guesses {
		if r == g {
			return true
		}
	}
	return false
}
