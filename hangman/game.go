package hangman

import "fmt"

// MaxGuesses the maximum number of guesses allowed
const MaxGuesses = 7

// Game tracks game state
type Game struct {
	Tally    *Tally
	wordList *WordList
	word     string
	guesses  []rune
}

// NewGame initiates a new hangman game
func NewGame(path string) (g *Game, err error) {
	var wl *WordList
	if wl, err = NewWordList(path); err != nil {
		return
	}
	g = &Game{wordList: wl, Tally: &Tally{}}
	g.Reset()

	return
}

// Reset a game and generates a new guess word
func (g *Game) Reset() {
	g.word = g.wordList.PickWord()
	g.guesses = []rune{}
	g.Tally.reset(MaxGuesses, g.word)
}

// Guess a letter in the game
func (g *Game) Guess(r rune) (*Tally, error) {
	if g.Tally.Status == Won || g.Tally.Status == Lost {
		return g.Tally, fmt.Errorf("this game is finished")
	}
	g.Tally.update(r)

	return g.Tally, nil
}
