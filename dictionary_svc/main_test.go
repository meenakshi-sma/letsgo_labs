package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/imhotepio/letsgo_labs/jsondic"
	"github.com/stretchr/testify/assert"
)

func TestNewWordHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/new_word?dic=artists", nil)

	w := httptest.NewRecorder()
	newWordHandler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	e := jsondic.Entry{}
	err := json.NewDecoder(w.Body).Decode(&e)
	assert.Nil(t, err)

	assert.Equal(t, "artists", e.Dictionary)
	assert.Equal(t, "assets", e.Location)
	assert.NotEqual(t, "", e.Word)
}
