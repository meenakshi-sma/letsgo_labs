package hangman

import (
	"fmt"
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

// CurrentTally returns the current game state
func (g *Game) CurrentTally() *Tally {
	return g.tally
}

// Guess a letter in the game
func (g *Game) Guess(r rune) (*Tally, error) {
	if g.tally.Status == Won || g.tally.Status == Lost {
		return g.tally, fmt.Errorf("game over")
	}

	if inRunes(g.guesses, r) {
		g.tally.Status = AlreadyGuessed
		return g.tally, nil
	}

	g.guesses = append(g.guesses, r)
	g.tally.Letters = lettersFromGuess(g.word, g.guesses)

	if inRunes([]rune(g.word), r) {
		g.tally.Status = CorrectGuess
		if !missingLetters(g.tally.Letters) {
			g.tally.Status = Won
		}
		return g.tally, nil
	}

	g.tally.TurnsLeft--
	g.tally.Status = IncorrectGuess
	if g.tally.TurnsLeft == 0 {
		g.tally.Status = Lost
		g.tally.Word = g.word
	}
	return g.tally, nil
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
