package hangman_test

import (
	"testing"

	"github.com/imhotepio/letsgo_labs/hangman"
	"github.com/stretchr/testify/assert"
)

func TestWordListOk(t *testing.T) {
	wl, err := hangman.NewWordList("./assets/test.txt")

	assert.Nil(t, err)
	assert.Equal(t, 3, wl.Len())
}

func TestWordListNotOk(t *testing.T) {
	_, err := hangman.NewWordList("./assets/testo.txt")
	assert.NotNil(t, err)
}

func TestWordPickWord(t *testing.T) {
	wl, err := hangman.NewWordList("./assets/test.txt")
	assert.Nil(t, err)

	w := wl.PickWord()

	assert.Equal(t, 4, len(w))
}
