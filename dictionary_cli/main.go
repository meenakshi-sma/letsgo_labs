package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/imhotepio/letsgo_labs/jsondic"
)

func main() {
	var (
		dic  = flag.String("dic", "words", "Dictionary name. Default words")
		host = flag.String("h", "localhost", "Dictionary service host. Default localhost")
		port = flag.Int("p", 4500, "Dictionary service port. Default 4500")
	)

	flag.Parse()

	e, err := dialADic(*dic, *host, *port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[DialADic] The word of the day is `%s\n", e.Word)
}

// ToRoman converts an arabic number to a roman glyph
func dialADic(dic, host string, port int) (jsondic.Entry, error) {
	return get(host, port, "new_word", map[string]string{"dic": dic})
}

func get(host string, port int, path string, args map[string]string) (jsondic.Entry, error) {
	var entry jsondic.Entry

	u, err := url.Parse(fmt.Sprintf("http://%s:%d/", host, port))
	if err != nil {
		return entry, err
	}
	u.Path = path

	vals := u.Query()
	for k, v := range args {
		vals.Set(k, v)
	}
	u.RawQuery = vals.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return entry, err
	}

	if resp.StatusCode != 200 {
		return entry, fmt.Errorf("Response not ok %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&entry)
	if err != nil {
		return entry, err
	}

	return entry, nil
}
