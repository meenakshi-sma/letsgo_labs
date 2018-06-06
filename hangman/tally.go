package hangman

const (
	// Started denotes a game in progress
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
		Word      string
	}
)

// NewTally initializes a new tally with a given number of turns
func newTally(maxTurns int, word string) *Tally {
	return &Tally{
		TurnsLeft: maxTurns,
		Letters:   lettersFromGuess(word, []rune{}),
	}
}
