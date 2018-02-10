package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const ServerURL = "http://localhost:9000/"

type response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
	URL    string `json:"url"`
}

func get(path string, args map[string]string) (string, error) {
	u, err := url.Parse(ServerURL)
	if err != nil {
		return "", err
	}
	u.Path = path

	vals := u.Query()
	for k, v := range args {
		vals.Set(k, v)
	}
	u.RawQuery = vals.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Response not ok %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body := response{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return "", err
	}

	return body.Result, nil
}

// ToRoman converts an arabic number to a roman glyph
func ToRoman(n int) (string, error) {
	return get("roman", map[string]string{"n": fmt.Sprintf("%d", n)})
}

// ToArabic converts a roman glyph to arabic
func ToArabic(g string) (int, error) {
	res, err := get("arabic", map[string]string{"g": g})
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res)
}

func main() {
	a := flag.Int("a", -1, "Provide an arabic number. Default 10")
	g := flag.String("g", "", "Provide a roman literal. Default X")

	flag.Parse()

	if *a != -1 {
		g, err := ToRoman(*a)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[DialARoman] Roman literal for %d is `%s\n", *a, g)
		return
	}

	if *g != "" {
		a, err := ToArabic(*g)
		if err != nil {
			panic(err)
		}
		fmt.Printf("[DialARoman] Arabic number for %s is `%d\n", *g, a)
	}
}
