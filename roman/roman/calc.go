// Package roman provides calculators for converting an arabic number
// to a roman glyph and vice versa
package roman

import "fmt"

var (
	limits = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	glyphs = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

// ToRoman converts an arabic number to a Roman glyph
func ToRoman(n int) (r string, e error) {
	if n > 3999 {
		r, e = "", fmt.Errorf("number is too large")
		return
	}

	for i, tick := range limits {
		for n >= tick {
			r += glyphs[i]
			n -= tick
		}
	}
	return
}

// ToArabic converts a roman glyph to arabic
func ToArabic(s string) (int, error) {
	res, chars := 0, []rune(s)
	for i, r := range s {
		i1, err := indexOfSymbol(r)
		if err != nil {
			return 0, err
		}

		if i+1 < len(s) {
			i2, err := indexOfSymbol(chars[i+1])
			if err != nil {
				return 0, err
			}
			if i1 > i2 {
				res -= limits[i1]
				continue
			}
		}
		res += limits[i1]
	}
	return res, nil
}

func indexOfSymbol(s rune) (int, error) {
	for i := range glyphs {
		if glyphs[i] == string(s) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("invalid roman glyph `%c", s)
}
