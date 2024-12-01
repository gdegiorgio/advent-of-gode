package main

import "testing"

type Input struct {
	a []int
	b []int
}

func TestResolve(t *testing.T) {
	cases := []struct {
		Description string
		Input       Input
		Distance    float64
	}{
		{
			Description: "Example Input",
			Input:       Input{a: []int{3, 4, 2, 1, 3, 3}, b: []int{4, 3, 5, 3, 9, 3}},
			Distance:    11,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			if got := resolve(c.Input.a, c.Input.b); got != c.Distance {
				t.Errorf("resolve(%v, %v) = %f, want %f", c.Input.a, c.Input.b, got, c.Distance)
			}
		})
	}
}
