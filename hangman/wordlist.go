package hangman

import (
	"bufio"
	"math/rand"
	"os"
)

// WordList tracks a collection of words
type WordList struct {
	words []string
	path  string
}

// NewWordList creates a new word list from a given list of words
func NewWordList(path string) (*WordList, error) {
	wl := &WordList{path: path}
	return wl, wl.load()
}

// Len of the collection of words in the list
func (w *WordList) Len() int {
	return len(w.words)
}

func (w *WordList) load() (err error) {
	var f *os.File
	if f, err = os.Open(w.path); err != nil {
		return
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		w.words = append(w.words, s.Text())
	}
	return
}

// PickWord selects a new word at random
func (w *WordList) PickWord() string {
	return w.words[rand.Intn(w.Len())]
}
