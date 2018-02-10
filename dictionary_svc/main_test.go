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
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/new_word?dic=artists", nil)
	)

	setUpRoutes()
	newWordHandler(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)

	e := jsondic.Entry{}
	err := json.NewDecoder(rr.Body).Decode(&e)
	assert.Nil(t, err)

	assert.Equal(t, "artists", e.Dictionary)
	assert.Equal(t, "assets", e.Location)
	assert.NotEqual(t, "", e.Word)
}

func TestNoMatchHandler(t *testing.T) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/", nil)
	)

	noMatchHandler(rr, r)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestNewWordHandlerNoDic(t *testing.T) {
	var (
		w      = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "http://example.com/new_word?dic=bozo", nil)
	)

	newWordHandler(w, req)
	assert.Equal(t, http.StatusExpectationFailed, w.Code)
}
