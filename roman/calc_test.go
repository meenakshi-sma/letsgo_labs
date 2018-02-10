package roman_test

import (
	"testing"

	"github.com/imhotepio/letsgo_labs/roman"
	"github.com/stretchr/testify/assert"
)

var useCases = []struct {
	arabic int
	glyph  string
}{
	{1, "I"},
	{2, "II"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{20, "XX"},
	{21, "XXI"},
	{25, "XXV"},
	{30, "XXX"},
	{35, "XXXV"},
	{40, "XL"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{2016, "MMXVI"},
	{3999, "MMMCMXCIX"},
	{0, ""},
}

func TestToRoman(t *testing.T) {
	for _, uc := range useCases {
		r, err := roman.ToRoman(uc.arabic)
		assert.Nil(t, err)

		assert.Equal(t, uc.glyph, r)
	}
}

func TestToRomanInvalid(t *testing.T) {
	_, err := roman.ToRoman(4000)

	assert.Errorf(t, err, "number is too large `%d", 3999)
}

func TestToArabic(t *testing.T) {
	for _, uc := range useCases {
		a, err := roman.ToArabic(uc.glyph)
		assert.Nil(t, err)

		assert.Equal(t, uc.arabic, a)
	}
}

func TestToArabicInvalid(t *testing.T) {
	for _, g := range []string{"IF", "FI"} {
		_, err := roman.ToArabic(g)

		assert.Errorf(t, err, "invalid roman glyph `%s", "F")
	}
}
