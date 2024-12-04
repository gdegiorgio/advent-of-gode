package main

import "testing"

func TestResolve(t *testing.T) {
	m := [][]string{
		{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
		{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
		{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
		{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
		{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
		{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
		{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
		{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
		{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
	}
	result := resolvePartOne(m)
	expected := 18

	if result != expected {
		t.Errorf("resolvePartOne(M)=%d; want %d", result, expected)
	}

}
func TestIsValidXMAS(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Expected    bool
	}{
		{Description: "Empty", Input: "", Expected: false},
		{Description: "Horizontal Value", Input: "XMAS", Expected: true},
		{Description: "Inverted", Input: "SAMX", Expected: true},
		{Description: "Invalid", Input: "XAMSX", Expected: false},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			actual := isValidXMAS(c.Input)
			if actual != c.Expected {
				t.Errorf("isValidXMAS(%q) == %t, want %t", c.Input, actual, c.Expected)
			}
		})
	}
}
