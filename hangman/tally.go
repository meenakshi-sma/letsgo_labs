// Package hangman
// Tally tracks the progression of guesses in the hangman game

package hangman

const (
	// Active indicates game is in progress
	Active GameStatus = iota
	// Won indicates the game is won
	Won
	// Lost indicates the game is lost
	Lost
	// AlreadyGuessed indicates a letter is already guessed
	AlreadyGuessed
)

type (
	// GameStatus tracks current game status
	GameStatus int

	// Tally tracks hangman outcome based on a guess
	Tally struct {
		TurnsLeft int
		guesses   []rune
		Letters   []rune
		Status    GameStatus
	}
)

// NewTally creates a new tally
func NewTally(turns int, wordLen int) *Tally {
	return &Tally{
		TurnsLeft: turns,
		Letters:   make([]rune, wordLen),
		Status:    Active,
	}
}

// Update a tally and recompute state
func (t *Tally) Update(w string, g rune) {
	if t.Status == Lost {
		return
	}

	if alreadyGuessed(t.guesses, g) {
		t.Status = AlreadyGuessed
		return
	}

	t.guesses = append(t.guesses, g)
	t.updateLetters(w)
	t.updateStatus(goodGuess(w, g))
}

func (t *Tally) updateLetters(w string) {
	for i, l := range w {
		if alreadyGuessed(t.guesses, l) {
			t.Letters[i] = l
		} else {
			t.Letters[i] = '_'
		}
	}
}

func (t *Tally) updateStatus(goodGuess bool) {
	if !missingLetters(t.Letters) {
		t.Status = Won
		return
	}

	if !goodGuess {
		t.TurnsLeft--
	}

	if t.TurnsLeft == 0 {
		t.Status = Lost
	} else {
		t.Status = Active
	}
}

// Supporting functions...

func missingLetters(letters []rune) bool {
	for _, l := range letters {
		if l == '_' {
			return true
		}
	}
	return false
}

func goodGuess(w string, g rune) bool {
	for _, l := range w {
		if l == g {
			return true
		}
	}
	return false
}

func alreadyGuessed(guesses []rune, l rune) bool {
	for _, g := range guesses {
		if g == l {
			return true
		}
	}
	return false
}
