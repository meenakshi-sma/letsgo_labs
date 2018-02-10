package jsondic_test

import (
	"testing"

	"github.com/imhotepio/letsgo_labs/jsondic"
	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	e := jsondic.Entry{
		Dictionary: "artists",
		Location:   "/tmp/assets",
		Word:       "bumblebeetuna",
	}
	r, err := jsondic.Encode(e)

	assert.Nil(t, err)
	assert.Equal(t, r, `{"dictionary":"artists","location":"/tmp/assets","random_word":"bumblebeetuna"}`+"\n")
}

func TestDecoding(t *testing.T) {
	raw := `{"dictionary":"artists","location":"/tmp/assets","random_word":"bumblebeetuna"}`

	e, err := jsondic.Decode(raw)

	assert.Nil(t, err)
	assert.Equal(t, "artists", e.Dictionary)
	assert.Equal(t, "/tmp/assets", e.Location)
	assert.Equal(t, "bumblebeetuna", e.Word)
}

func TestDecodingFail(t *testing.T) {
	raw := `fred`

	_, err := jsondic.Decode(raw)
	assert.NotNil(t, err)
}
