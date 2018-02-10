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

	err := json.NewEncoder(&buff).Encode(e)
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}

// Decode hydrates an entry from JSON
func Decode(raw string) (Entry, error) {
	var e Entry
	err := json.NewDecoder(bytes.NewReader([]byte(raw))).Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}
