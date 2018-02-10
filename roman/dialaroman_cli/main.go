package main

const ServerURL = "http://localhost:9000/"

type response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
	URL    string `json:"url"`
}

// ToRoman converts an arabic number to a roman glyph
func ToRoman(n int) (string, error) {
	// YOUR_CODE...
}

// ToArabic converts a roman glyph to arabic
func ToArabic(g string) (int, error) {
	// YOUR_CODE...
}

func main() {
	// YOUR_CODE...
}

func get(path string, args map[string]string) (string, error) {
	// YOUR_CODE...
}
