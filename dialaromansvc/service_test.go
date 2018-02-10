package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArabicHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/arabic?g=X", nil)

	w := httptest.NewRecorder()
	toArabicHandler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	resp := response{}
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.Nil(t, err)

	assert.Equal(t, "10", resp.Result)
}

func TestRomainHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://example.com/roman?n=10", nil)

	w := httptest.NewRecorder()
	toRomanHandler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	resp := response{}
	err := json.NewDecoder(w.Body).Decode(&resp)
	assert.Nil(t, err)

	assert.Equal(t, "X", resp.Result)
}
