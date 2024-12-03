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
		Similarity  int
	}{
		{
			Description: "Example Input",
			Input:       Input{a: []int{3, 4, 2, 1, 3, 3}, b: []int{4, 3, 5, 3, 9, 3}},
			Distance:    11,
			Similarity:  31,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			distance, similarity := resolve(c.Input.a, c.Input.b)
			if distance != c.Distance {
				t.Errorf("resolve(%v, %v) = %.0f ; want distance %.0f", c.Input.a, c.Input.b, distance, c.Distance)
			}
			if similarity != c.Similarity {
				t.Errorf("resolve(%v, %v) = %d ; want similarity %d", c.Input.a, c.Input.b, similarity, c.Similarity)
			}
		})
	}
}
