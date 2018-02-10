package main

import (
	"net/http"
)

const httpPort = ":4500"

func main() {
	setupRoutes()
	// YOUR_CODE...
}

func setupRoutes() {
	// YOUR_CODE...
}

func newWordHandler(w http.ResponseWriter, r *http.Request) {
	// YOUR_CODE...
}

func noMatchHandler(w http.ResponseWriter, r *http.Request) {
	// YOUR_CODE...
}
