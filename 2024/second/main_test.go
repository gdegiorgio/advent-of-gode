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

func TestIsSafe(t *testing.T) {
	cases := []struct {
		Description string
		Input       []int
		Safe        int
	}{
		{
			Description: "Example Input 1",
			Input:       []int{1, 3, 6, 7, 9},
			Safe:        1,
		},
		{
			Description: "Example Input 2",
			Input:       []int{8, 6, 4, 4, 1},
			Safe:        1,
		},
		{
			Description: "Example Input 3",
			Input:       []int{8, 6, 4, 4, 4, 1},
			Safe:        0,
		},
		{
			Description: "Example Input 4",
			Input:       []int{8, 6, 4, 2, 7},
			Safe:        1,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			safe := isSafe(c.Input, true)
			if safe != c.Safe {
				t.Errorf("isSafe(%v) = %d ; want %d", c.Input, safe, c.Safe)
			}
		})
	}
}

func TestRemoveItem(t *testing.T) {
	arr := []int{1, 2, 4}
	expected := []int{1, 4}
	actual := removeItem(arr, 1)
	if len(actual) != len(arr)-1 {
		t.Errorf("removeItem(%v, 1) = %v ; want %v", arr, actual, expected)
	}

}
