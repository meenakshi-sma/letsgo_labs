package hangman_test

import (
	"fmt"
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

func TestExtraGuesses(t *testing.T) {
	g := hangman.NewGame("test")

	for _, l := range "abcdfghi" {
		g.Guess(l)
	}

	assert.Equal(t, 0, g.Tally.TurnsLeft)
	assert.Equal(t, hangman.Lost, g.Tally.Status)
}

func ExampleNewGame() {
	g := hangman.NewGame("test")
	fmt.Println(string(g.Tally.Letters), g.Tally.TurnsLeft, g.Tally.Status)
	// Output:
	// ____ 7 0
}

func ExampleGame_Guess_good() {
	g := hangman.NewGame("test")
	g.Guess('t')
	fmt.Println(string(g.Tally.Letters), g.Tally.TurnsLeft, g.Tally.Status)
	// Output:
	// t__t 7 0
}

func ExampleGame_Guess_bad() {
	g := hangman.NewGame("test")
	g.Guess('z')
	fmt.Println(string(g.Tally.Letters), g.Tally.TurnsLeft, g.Tally.Status)
	// Output:
	// ____ 6 0
}

func ExampleGame_Guess_win() {
	g := hangman.NewGame("test")
	for _, l := range []rune("tes") {
		g.Guess(l)
	}
	fmt.Println(string(g.Tally.Letters), g.Tally.TurnsLeft, g.Tally.Status)
	// Output:
	// test 7 1
}

func ExampleGame_Guess_loose() {
	g := hangman.NewGame("test")
	for _, l := range []rune("abcdfgh") {
		g.Guess(l)
	}
	fmt.Println(string(g.Tally.Letters), g.Tally.TurnsLeft, g.Tally.Status)
	// Output:
	// ____ 0 2
}

func ExampleGame_Guess_alreadyGuessed() {
	g := hangman.NewGame("test")
	for _, l := range []rune("tee") {
		g.Guess(l)
	}
	fmt.Println(string(g.Tally.Letters), g.Tally.TurnsLeft, g.Tally.Status)
	// Output:
	// te_t 7 3
}
