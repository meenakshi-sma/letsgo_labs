// © 2018 Imhotep Software LLC. All rights reserved.

// Package hangman
// Implements a hangman game. The game tracks a word to be guessed
// and a tally to track game state and progress.
//
// Playing
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

// Guess a letter and computes a corresponding tally
func (g *Game) Guess(guess rune) {
	g.Tally.Update(g.word, guess)
}
