package romannumerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	tests := []struct {
		Description string
		Number      int
		Want        string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Number)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}
