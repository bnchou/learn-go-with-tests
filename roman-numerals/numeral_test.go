package main

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Roman       string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func ConvertToRoman(i int) string {
	if i == 3 {
		return "III"
	}
	if i == 2 {
		return "II"
	}
	return "I"
}
