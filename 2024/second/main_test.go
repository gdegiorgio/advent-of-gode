package main

import "testing"

func TestResolve(t *testing.T) {
	cases := []struct {
		Description string
		Input       [][]int
		Safe        int
	}{
		{
			Description: "Example Input",
			Input: [][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			Safe: 2,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			safe := resolve(c.Input)
			if safe != c.Safe {
				t.Errorf("resolve(%v) = %d ; want %d", c.Input, safe, c.Safe)
			}
		})
	}
}
