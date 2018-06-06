package hangman

import (
	"math/rand"
	"time"
)

// MaxGuesses the maximum number of guesses allowed
const MaxGuesses = 7

// Game tracks game state
type Game struct {
	wordList *WordList
	guesses  []rune
	word     string
	tally    *Tally
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewGame initiates a new hangman game
func NewGame(path string) (g *Game, err error) {
	var wl *WordList
	if wl, err = NewWordList(path); err != nil {
		return
	}
	word := wl.PickWord()
	g = &Game{wordList: wl, word: word, tally: newTally(MaxGuesses, word)}
	return
}

// Guess a letter in the game
func (g *Game) Guess(r rune) (*Tally, error) {
	// YOUR_CODE...
}

func lettersFromGuess(w string, guesses []rune) []rune {
	res := make([]rune, len(w))
	for i, r := range w {
		if inRunes(guesses, r) {
			res[i] = r
		} else {
			res[i] = '-'
		}
	}
	return res
}

func inRunes(rr []rune, g rune) bool {
	for _, r := range rr {
		if r == g {
			return true
		}
	}
	return false
}

func missingLetters(ll []rune) bool {
	return inRunes(ll, '-')
}
