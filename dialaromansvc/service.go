package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/imhotepio/letsgo_labs/roman"
)

const httpPort = ":9000"

type response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
	URL    string `json:"url"`
}

func toRomanHandler(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	glyph, err := roman.ToRoman(number)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}
	resp := response{
		Status: http.StatusOK,
		Result: glyph,
		URL:    fmt.Sprintf("http://%s/arabic?g=%s", r.Host, glyph),
	}

	buff := new(bytes.Buffer)
	err = json.NewEncoder(buff).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, buff.String())
	log.Printf("[%d] %s", http.StatusOK, r.URL)
}

func toArabicHandler(w http.ResponseWriter, r *http.Request) {
	n, err := roman.ToArabic(r.URL.Query().Get("g"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := response{
		Status: http.StatusOK,
		Result: fmt.Sprintf("%d", n),
		URL:    fmt.Sprintf("http://%s/roman?n=%d", r.Host, n),
	}

	buff := new(bytes.Buffer)
	err = json.NewEncoder(buff).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, buff.String())
	log.Printf("[%d] %s", http.StatusOK, r.URL)
}

func noMatchHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("No matching routes! %s", r.URL), http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/roman", toRomanHandler)
	http.HandleFunc("/arabic", toArabicHandler)
	http.HandleFunc("/", noMatchHandler)

	log.Printf("DialARoman is listening on port [%s]", httpPort)
	http.ListenAndServe(httpPort, nil)
}
