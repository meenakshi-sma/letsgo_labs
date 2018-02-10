package dictionary_test

import (
	"fmt"
	"testing"

	"github.com/imhotepio/letsgo_labs/dictionary"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	l, err := dictionary.LoadDefault("musicians")

	assert.Nil(t, err)
	assert.Equal(t, 5, len(l))
}

func TestPicker(t *testing.T) {
	l, err := dictionary.LoadDefault("musicians")
	w := dictionary.Pick(l)

	assert.Nil(t, err)
	assert.NotEqual(t, "", w)
}

func TestLoadMissing(t *testing.T) {
	_, err := dictionary.LoadDefault("actors")

	assert.EqualError(t, err, "unable to load dictionary `assets/actors.txt")
}

func TestLoadCustomFail(t *testing.T) {
	_, err := dictionary.Load("zorg", "actors")
	fmt.Println(err)
	assert.EqualError(t, err, "unable to load dictionary `zorg/actors.txt")
}
