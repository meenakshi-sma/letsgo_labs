// © 2018 Imhotep Software LLC. All rights reserved.

// Package jsondic provides for JSON encoding/decoding of a dictionary entry.
package jsondic

import (
	"bytes"
	"encoding/json"
)

// Entry from a dictionary load
type Entry struct {
	Dictionary string `json:"dictionary"`
	Location   string `json:"location"`
	Word       string `json:"random_word"`
}

// Encode an entry into JSON
func Encode(e Entry) (string, error) {
	var buff bytes.Buffer
	if err := json.NewEncoder(&buff).Encode(e); err != nil {
		return "", err
	}
	return buff.String(), nil
}

// Decode hydrates an entry from JSON
func Decode(raw string) (e Entry, err error) {
	err = json.NewDecoder(bytes.NewReader([]byte(raw))).Decode(&e)
	return
}
