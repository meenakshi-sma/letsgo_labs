package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/imhotepio/letsgo_labs/dictionary"
	"github.com/imhotepio/letsgo_labs/jsondic"
)

const httpPort = ":4500"

func main() {
	setUpRoutes()

	log.Printf("Private Dic is listening [%s]", httpPort)
	http.ListenAndServe(httpPort, nil)
}

func setUpRoutes() {
	http.HandleFunc("/new_word", newWordHandler)
	http.HandleFunc("/", noMatchHandler)
}

func newWordHandler(w http.ResponseWriter, r *http.Request) {
	dic := r.URL.Query().Get("dic")

	wl, err := dictionary.LoadDefault(dic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	buff, err := jsondic.Encode(jsondic.Entry{
		Dictionary: dic,
		Location:   "assets",
		Word:       dictionary.Pick(wl),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, buff)
	log.Printf("[%d] %s", http.StatusOK, r.URL)
}

func noMatchHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("No matching routes! %s", r.URL), http.StatusBadRequest)
}
