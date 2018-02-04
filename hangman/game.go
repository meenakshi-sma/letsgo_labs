// Package Hangman
// Plays a hangman game using a give word

package hangman

// Game tracks hangman game state and tally
type Game struct {
	word  string
	Tally *Tally
}

// NewGame initializes a new hangman game
func NewGame(w string) *Game {
	return &Game{word: w, Tally: NewTally(7, len(w))}
}

// Guess a new letter and computes a new game tally
func (g *Game) Guess(guess rune) {
	g.Tally.Update(g.word, guess)
}
